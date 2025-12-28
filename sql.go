package snowflake

import (
	"database/sql/driver"
	"fmt"
)

func (id *ID) Scan(v interface{}) error {
	switch v := v.(type) {
	case nil:
		return nil
	case int64:
		*id = ID(v)
	case string:
		if v == "" {
			return nil
		}
		u, err := ParseString(v)
		if err != nil {
			return fmt.Errorf("scan: %v", err)
		}
		*id = u
	case []byte:
		if len(v) == 0 {
			return nil
		}
		return id.Scan(string(v))
	default:
		return fmt.Errorf("scan: unable to scan type %T into Snowflake ID", v)
	}
	return nil
}

func (id ID) Value() (driver.Value, error) {
	return id.Int64(), nil
}
