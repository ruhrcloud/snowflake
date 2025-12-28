package snowflake

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"sync"
	"time"
)

var (
	ErrNodeIDInvalid  = errors.New("node ID must be between 0 and 1023")
	ErrClockBackwards = errors.New("clock moved backwards")
)

type Node struct {
	mu    sync.Mutex
	epoch int64
	node  int64

	time int64
	step int64
}

func New(id int64) (*Node, error) {
	// If id is -1, a random 10-bit instance ID is generated
	if id == -1 {
		var byt [8]byte
		_, _ = rand.Read(byt[:])
		id = int64(binary.BigEndian.Uint64(byt[:]) & uint64(MaxNode))
	}

	if id < 0 || id > MaxNode {
		return nil, ErrNodeIDInvalid
	}

	return &Node{
		epoch: BaseEpoch,
		node:  id,
	}, nil
}

func NewFromParts(datacenterID, workerID int64) (*Node, error) {
	if datacenterID < 0 || datacenterID > MaxDatacenter {
		return nil, errors.New("datacenter ID must be between 0 and 31")
	}
	if workerID < 0 || workerID > MaxWorker {
		return nil, errors.New("worker ID must be between 0 and 31")
	}

	id := (datacenterID << WorkerBits) | workerID
	return New(id)
}

func (n *Node) SetEpoch(epoch int64) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.epoch = epoch
}

func (n *Node) Generate() (ID, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixMilli()

	if now < n.epoch {
		return 0, errors.New("current time is before epoch")
	}

	if now < n.time {
		// We wait for the clock to catch up if the drift is within 5ms
		drift := n.time - now
		if drift <= 5 {
			time.Sleep(time.Duration(drift) * time.Millisecond)
			now = time.Now().UnixMilli()
		}

		if now < n.time {
			return 0, ErrClockBackwards
		}
	}

	if now == n.time {
		n.step = (n.step + 1) & MaxStep
		if n.step == 0 {
			// We increase the timestamp by 1ms to avoid sequence overflow
			now = n.time + 1
		}
	} else if now > n.time {
		n.step = 0
	}

	n.time = now
	r := ID((now-n.epoch)<<TimestampShift |
		(n.step << StepShift) |
		(n.node << NodeShift))

	return r, nil
}
