package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/update-service/entity"
)

type updateRepository struct {
	db *sqlx.DB
}

type UpdateRepository interface {
	Create(entity.Student) error
}

func NewUpdateRepository(db *sqlx.DB) UpdateRepository {
	return &updateRepository{
		db: db,
	}
}

func (r *updateRepository) Create(student entity.Student) error {
	if _, err := r.db.Query("INSERT INTO student(npm, name) VALUES($1, $2)", student.NPM, student.Name); err != nil {
		switch e := err.(type) {
		case *pq.Error:
			switch e.Code {
			case "23505": // unique constraint
				return errors.Wrap(entity.ErrInvalidNPM, "there exist student with npm "+student.NPM)
			}
		}
	}

	return nil
}
