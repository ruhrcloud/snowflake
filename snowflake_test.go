package snowflake

import (
	"encoding/json"
	"testing"
)

func TestParseString(t *testing.T) {
	id, err := ParseString("123456789")
	if err != nil {
		t.Fatalf("Failed to parse string: %v", err)
	}
	if id != 123456789 {
		t.Errorf("Expected 123456789, got %v", id)
	}
}

func TestJSON(t *testing.T) {
	id := ID(123456789)
	data, err := json.Marshal(id)
	if err != nil {
		t.Fatalf("Failed to marshal ID: %v", err)
	}

	var id2 ID
	if err := json.Unmarshal(data, &id2); err != nil {
		t.Fatalf("Failed to unmarshal ID: %v", err)
	}

	if id != id2 {
		t.Errorf("Unmarshaled ID doesn't match: %v != %v", id, id2)
	}
}

func TestDecompose(t *testing.T) {
	id := ID((1 << 22) | (2 << 10) | 3)
	parts := id.Decompose()

	if parts.Timestamp != BaseEpoch+1 {
		t.Errorf("Expected timestamp %v, got %v", BaseEpoch+1, parts.Timestamp)
	}
	if parts.Step != 2 {
		t.Errorf("Expected step 2, got %v", parts.Step)
	}
	if parts.Node != 3 {
		t.Errorf("Expected node 3, got %v", parts.Node)
	}

	if id.Timestamp() != BaseEpoch+1 {
		t.Errorf("Expected timestamp %v, got %v", BaseEpoch+1, id.Timestamp())
	}
	if id.Step() != 2 {
		t.Errorf("Expected step 2, got %v", id.Step())
	}
	if id.Node() != 3 {
		t.Errorf("Expected node 3, got %v", id.Node())
	}
}

func TestEqual(t *testing.T) {
	id1 := ID(12345)
	id2 := ID(12345)
	id3 := ID(67890)

	if !id1.Equal(id2) {
		t.Error("expected 12345 to equal 12345")
	}
	if id1.Equal(id3) {
		t.Error("expected 12345 to not equal 67890")
	}
}

func TestCompare(t *testing.T) {
	id1 := ID(100)
	id2 := ID(200)

	if id1.Compare(id1) != 0 {
		t.Error("expected 100 to compare 0 with 100")
	}
	if id1.Compare(id2) != -1 {
		t.Error("expected 100 to compare -1 with 200")
	}
	if id2.Compare(id1) != 1 {
		t.Error("expected 200 to compare 1 with 100")
	}
}
