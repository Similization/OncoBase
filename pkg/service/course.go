package services

import (
	"med/pkg/model"
	"med/pkg/repository"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) CreateCourse(course model.Course) (model.Course, error) {
	return s.repo.CreateCourse(course)
}
func (s *CourseService) GetCourseById(id string) (model.Course, error) {
	return s.repo.GetCourseById(id)
}
func (s *CourseService) GetCourseList() ([]model.Course, error) {
	return s.repo.GetCourseList()
}
func (s *CourseService) UpdateCourse(course model.Course) (model.Course, error) {
	return s.repo.UpdateCourse(course)
}
func (s *CourseService) DeleteCourse(id string) error {
	return s.repo.DeleteCourse(id)
}
