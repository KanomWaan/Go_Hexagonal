package repository

import "errors"

type courseRepositoryMock struct {
	customers []Course
}

func NewCourseRepositoryMock() courseRepositoryMock {
	courses := []Course{
		{ID: 99, Name: "mock course", Price: 999, Instructor: "mock instructor"},
		{ID: 100, Name: "mock course 2", Price: 999, Instructor: "mock instructor 2"},
	}
	return courseRepositoryMock{customers: courses}
}

func (r courseRepositoryMock) GetAll() ([]Course, error) {
	return r.customers, nil
}

func (r courseRepositoryMock) GetByID(id int) (*Course, error) {
	for _, course := range r.customers {
		if course.ID == id {
			return &course, nil
		}
	}
	return nil, errors.New("course not found")
}
