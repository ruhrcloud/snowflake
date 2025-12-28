package snowflake

import "testing"

func TestSQLScan(t *testing.T) {
	var id ID
	if err := id.Scan(int64(12345)); err != nil {
		t.Fatalf("failed to scan int64: %v", err)
	}
	if id != 12345 {
		t.Errorf("expected 12345, got %v", id)
	}

	if err := id.Scan("67890"); err != nil {
		t.Fatalf("failed to scan string: %v", err)
	}
	if id != 67890 {
		t.Errorf("expected 67890, got %v", id)
	}

	if err := id.Scan([]byte("13579")); err != nil {
		t.Fatalf("failed to scan []byte: %v", err)
	}
	if id != 13579 {
		t.Errorf("expected 13579, got %v", id)
	}

	if err := id.Scan(nil); err != nil {
		t.Fatalf("failed to scan nil: %v", err)
	}

	if err := id.Scan(true); err == nil {
		t.Error("expected error scanning bool")
	}
}

func TestSQLValue(t *testing.T) {
	id := ID(12345)
	val, _ := id.Value()
	if val != int64(12345) {
		t.Errorf("expected int64(12345), got %v", val)
	}
}
