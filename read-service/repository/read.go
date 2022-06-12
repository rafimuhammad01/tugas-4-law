package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rafimuhammad01/read-service/entity"
)

type readRepository struct {
	db *sqlx.DB
}

type ReadRepository interface {
	Get(string) (*entity.Student, error)
}

func NewReadRepository(db *sqlx.DB) ReadRepository {
	return &readRepository{
		db: db,
	}
}

func (r *readRepository) Get(npm string) (*entity.Student, error) {
	var studentInfo entity.Student

	if err := r.db.Get(&studentInfo, "SELECT id, npm, name FROM student WHERE npm = $1", npm); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(entity.ErrStudentNotFound, "student with npm "+npm+" is not exist")
		}

		return nil, err
	}

	return &studentInfo, nil
}
