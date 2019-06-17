package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
)

func main() {
	var demo struct {
		DateOfBirth NullTime `json:"date_of_birth" db:"date_of_birth"` // date_of_birth
	}

	//demo.DateOfBirth = NullTime{Time: time.Now(), Valid: true}
	jsonBuffer, err := json.Marshal(&demo)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", demo, string(jsonBuffer))

	var newDemo struct {
		DateOfBirth NullTime `json:"date_of_birth" db:"date_of_birth"` // date_of_birth
	}
	err = json.Unmarshal(jsonBuffer, &newDemo)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", newDemo)
}

var nullBytes = []byte(`""`)

type NullTime pq.NullTime

// MarshalJSON is a custom marshaler for Time
func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nullBytes, nil
	}
	return json.Marshal(t.Time)
}

// UnmarshalJSON is a custom unmarshaler for Time
func (t NullTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		t.Valid = false
		return nil
	}

	var tm time.Time
	err := json.Unmarshal(data, &tm)
	if err != nil {
		return err
	}

	t.Time = tm
	return nil
}
