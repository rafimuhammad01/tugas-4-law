package entity

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidNPM      = errors.New("invalid npm")
	ErrInvalidName     = errors.New("invalid name")
	ErrStudentNotFound = errors.New("student not found")
	ErrInvalidTrxID    = errors.New("invalid transaction id")
)

type Student struct {
	ID   int    `json:"id"`
	NPM  string `json:"npm"`
	Name string `json:"name"`
}

func (s Student) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Student) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	return nil
}
