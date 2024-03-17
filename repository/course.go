package repository

type Course struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Price      int    `db:"price"`
	Instructor string `db:"instructor"`
}

type CourseRepository interface {
	GetAll() ([]Course, error)
	GetByID(id int) (*Course, error)
} 