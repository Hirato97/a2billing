package model

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (s NullInt64) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.Int64)
	}
	return []byte(`null`), nil
}
