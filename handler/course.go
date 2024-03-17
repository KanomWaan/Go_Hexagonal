package handler

import (
	"encoding/json"
	"fmt"
	"hexagonal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type courseHandler struct {
	courseService service.CourseService
}

func NewCourseHandler(courseService service.CourseService) courseHandler {
	return courseHandler{courseService: courseService}
}

func (h courseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.courseService.GetCourses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func (h courseHandler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	courseID, _ := strconv.Atoi(mux.Vars(r)["id"])
	course, err := h.courseService.GetCourseByID(courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}
