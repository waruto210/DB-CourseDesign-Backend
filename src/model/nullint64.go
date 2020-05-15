package model

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 struct {
	sql.NullInt64
}

func (v *NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	var i *int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		v.Valid = true
		v.Int64 = *i
	} else {
		v.Valid = false
	}
	return nil
}
