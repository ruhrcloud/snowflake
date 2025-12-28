package snowflake

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestMarshalText(t *testing.T) {
	id := ID(12345)
	text, _ := id.MarshalText()
	if string(text) != "12345" {
		t.Errorf("expected \"12345\", got %q", text)
	}

	var id2 ID
	if err := id2.UnmarshalText(text); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if id != id2 {
		t.Errorf("expected %v, got %v", id, id2)
	}
}

func TestMarshalBinary(t *testing.T) {
	id := ID(12345)
	bin, _ := id.MarshalBinary()
	if len(bin) != 8 {
		t.Errorf("expected 8 bytes, got %d", len(bin))
	}

	var id2 ID
	if err := id2.UnmarshalBinary(bin); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if id != id2 {
		t.Errorf("expected %v, got %v", id, id2)
	}

	if err := id2.UnmarshalBinary([]byte{1, 2, 3}); err == nil {
		t.Error("expected error for 3 bytes")
	}
}

func TestMarshalJSON(t *testing.T) {
	id := ID(12345)
	data, _ := json.Marshal(id)
	if !bytes.Equal(data, []byte(`"12345"`)) {
		t.Errorf("expected \"12345\", got %s", data)
	}

	var id2 ID
	if err := json.Unmarshal(data, &id2); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if id != id2 {
		t.Errorf("expected %v, got %v", id, id2)
	}
}
