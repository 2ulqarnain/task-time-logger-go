package structs

import (
	"encoding/json"
	"time"
)

// NullTime represents a nullable time value
type NullTime time.Time

// MarshalJSON implements custom JSON marshaling
func (nt NullTime) MarshalJSON() ([]byte, error) {
	t := time.Time(nt)
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t)
}

// GobEncode implements the gob.GobEncoder interface
func (nt NullTime) GobEncode() ([]byte, error) {
	t := time.Time(nt)
	return t.GobEncode()
}

// GobDecode implements the gob.GobDecoder interface
func (nt *NullTime) GobDecode(data []byte) error {
	var t time.Time
	if err := t.GobDecode(data); err != nil {
		return err
	}
	*nt = NullTime(t)
	return nil
}
