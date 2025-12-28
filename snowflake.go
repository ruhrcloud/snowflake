package snowflake

import "strconv"

type ID uint64

const (
	TimestampBits = uint8(41)
	NodeBits      = uint8(10)
	StepBits      = uint8(12)

	DatacenterBits = uint8(5)
	WorkerBits     = uint8(5)

	MaxNode = int64(-1 ^ (-1 << NodeBits))
	MaxStep = int64(-1 ^ (-1 << StepBits))

	MaxDatacenter = int64(-1 ^ (-1 << DatacenterBits))
	MaxWorker     = int64(-1 ^ (-1 << WorkerBits))

	NodeShift      = uint8(0)
	StepShift      = NodeBits
	TimestampShift = NodeBits + StepBits
)

// TODO
// This is currently based on Twitters epoch
const BaseEpoch = int64(1288834974657)

func (f ID) Int64() int64 {
	return int64(f)
}

func (f ID) Timestamp() int64 {
	return (int64(f) >> TimestampShift) + BaseEpoch
}

func (f ID) Node() int64 {
	return int64(f) & MaxNode
}

func (f ID) Step() int64 {
	return (int64(f) >> StepShift) & MaxStep
}

func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

func (id ID) Equal(other ID) bool {
	return id == other
}

func (id ID) Compare(other ID) int {
	if id == other {
		return 0
	}
	if id < other {
		return -1
	}
	return 1
}

func ParseString(s string) (ID, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(i), nil
}

type Decomposed struct {
	Timestamp int64
	Node      int64
	Step      int64
}

func (f ID) Decompose() Decomposed {
	return Decomposed{
		Timestamp: (int64(f) >> TimestampShift) + BaseEpoch,
		Step:      (int64(f) >> StepShift) & MaxStep,
		Node:      int64(f) & MaxNode,
	}
}
