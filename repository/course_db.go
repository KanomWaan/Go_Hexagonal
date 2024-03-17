package repository

import (
	"github.com/jmoiron/sqlx"
)

type courseRepoDB struct {
	db *sqlx.DB
}

func NewCourseRepoDB(db *sqlx.DB) courseRepoDB {
	return courseRepoDB{db: db}
}

func (c courseRepoDB) GetAll() ([]Course, error) {
	course := []Course{}
	query := "SELECT id,name,price,instructor FROM courses"
	err := c.db.Select(&course, query)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c courseRepoDB) GetByID(id int) (*Course, error) {
	course := Course{}
	query := "SELECT id,name,price,instructor FROM courses WHERE id = ?"
	err := c.db.Get(&course, query, id)
	if err != nil {
		return nil, err
	}
	return &course, nil
}
