package service

type CourseResponse struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Instructor string `json:"instructor"`
}


type CourseService interface {
	GetCourses() ([]CourseResponse, error)
	GetCourseByID(id int) (*CourseResponse, error)
}