package snowflake

import (
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	if _, err := New(1023); err != nil {
		t.Fatalf("failed to create node: %v", err)
	}
	if _, err := New(1024); err == nil {
		t.Fatal("expected error for node ID 1024")
	}

	n, err := New(-1)
	if err != nil {
		t.Fatalf("failed to create random node: %v", err)
	}
	if n.node < 0 || n.node > MaxNode {
		t.Errorf("random node ID %d out of range", n.node)
	}
}

func TestGenerate(t *testing.T) {
	n, _ := New(1)
	id1, _ := n.Generate()
	id2, _ := n.Generate()

	if id1 == id2 {
		t.Errorf("identical IDs: %v", id1)
	}
	if id1 > id2 {
		t.Errorf("IDs not increasing: %v > %v", id1, id2)
	}
}

func TestUniqueness(t *testing.T) {
	n, _ := New(1)
	count := 100000
	ids := make(chan ID, count)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				id, _ := n.Generate()
				ids <- id
			}
		}()
	}

	wg.Wait()
	close(ids)

	uids := make(map[ID]struct{})
	for id := range ids {
		if _, ok := uids[id]; ok {
			t.Errorf("duplicate ID: %v", id)
		}
		uids[id] = struct{}{}
	}
}

func TestEpochError(t *testing.T) {
	n, _ := New(1)
	n.SetEpoch(time.Now().UnixMilli() + 100000)
	if _, err := n.Generate(); err == nil {
		t.Error("expected error for future epoch")
	}
}

func TestNewFromParts(t *testing.T) {
	n, err := NewFromParts(10, 20)
	if err != nil {
		t.Fatalf("failed to create node: %v", err)
	}

	id, _ := n.Generate()
	if id.Node() != 340 {
		t.Errorf("expected node ID 340, got %v", id.Node())
	}
}

func TestClockDrift(t *testing.T) {
	n, _ := New(1)
	n.Generate()

	n.mu.Lock()
	n.time += 1000
	n.mu.Unlock()

	if _, err := n.Generate(); err != ErrClockBackwards {
		t.Errorf("expected ErrClockBackwards, got %v", err)
	}
}

func TestOverflow(t *testing.T) {
	n, _ := New(1)
	n.mu.Lock()
	n.time = time.Now().UnixMilli()
	n.step = MaxStep
	n.mu.Unlock()

	next := n.time + 1
	id, err := n.Generate()
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	if id.Timestamp() != next {
		t.Errorf("expected overflowed timestamp %v, got %v", next, id.Timestamp())
	}
	if id.Step() != 0 {
		t.Errorf("expected step 0, got %v", id.Step())
	}
}
