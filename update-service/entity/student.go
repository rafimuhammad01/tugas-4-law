package entity

import (
	"github.com/pkg/errors"
	"strconv"
)

var (
	ErrInvalidNPM  = errors.New("invalid npm")
	ErrInvalidName = errors.New("invalid name")
)

type Student struct {
	ID   int    `json:"id"`
	NPM  string `json:"npm"`
	Name string `json:"name"`
}

func (s *Student) Validate() error {
	if s.NPM == "" {
		return errors.Wrap(ErrInvalidNPM, "npm should not be empty")
	}

	if s.Name == "" {
		return errors.Wrap(ErrInvalidName, "name should not be empty")
	}

	if _, err := strconv.Atoi(s.NPM); err != nil {
		return errors.Wrap(ErrInvalidNPM, "npm should only contains number")
	}

	return nil
}
