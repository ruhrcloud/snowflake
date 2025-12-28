package snowflake

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

func (id ID) MarshalText() ([]byte, error) {
	return []byte(id.String()), nil
}

func (id *ID) UnmarshalText(data []byte) error {
	u, err := ParseString(string(data))
	if err != nil {
		return err
	}
	*id = u
	return nil
}

func (id ID) MarshalBinary() ([]byte, error) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(id))
	return b, nil
}

func (id *ID) UnmarshalBinary(data []byte) error {
	if len(data) != 8 {
		return fmt.Errorf("invalid Snowflake ID (got %d bytes)", len(data))
	}
	*id = ID(binary.BigEndian.Uint64(data))
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + id.String() + `"`), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	u, err := ParseString(s)
	if err != nil {
		return err
	}
	*id = u
	return nil
}
