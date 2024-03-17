package service

import (
	"errors"
	"hexagonal/repository"
)

type courseService struct {
	courseRepo repository.CourseRepository
}

func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return courseService{courseRepo: courseRepo}
}

func (s courseService) GetCourses() ([]CourseResponse, error) {
	courses, err := s.courseRepo.GetAll()
	if err != nil {
		return nil, err
	}

	coursesResponse := []CourseResponse{}
	for _, course := range courses {
		coursesResponse = append(coursesResponse, CourseResponse{
			Name:       course.Name,
			Price:      course.Price,
			Instructor: course.Instructor,
		})
	}
	return coursesResponse, nil
}

func (s courseService) GetCourseByID(id int) (*CourseResponse, error) {
	course, err := s.courseRepo.GetByID(id)
	if err != nil {

		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("course not found")
		}
		return nil, err
	}

	courseResponse := CourseResponse{
		Name:       course.Name,
		Price:      course.Price,
		Instructor: course.Instructor,
	}
	return &courseResponse, nil
}
