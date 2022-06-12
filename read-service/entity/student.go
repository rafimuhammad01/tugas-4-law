package entity

import "errors"

var (
	ErrStudentNotFound = errors.New("student not found")
)

type Student struct {
	ID   int    `json:"id"`
	NPM  string `json:"npm"`
	Name string `json:"name"`
}
