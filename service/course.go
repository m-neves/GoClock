package service

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/repository"
)

type CourseServiceInterface interface {
	Create(course *data.Course, userId int) error
	FindAll(userId int) ([]*data.Course, error)
}

type courseService struct {
	cr repository.CourseRepositoryInterface
}

func NewCourseService(repository repository.CourseRepositoryInterface) CourseServiceInterface {
	return &courseService{repository}
}

func (cs *courseService) Create(course *data.Course, userId int) error {
	return cs.cr.Create(course, userId)
}

func (cs *courseService) FindAll(userId int) ([]*data.Course, error) {
	return cs.cr.FindAll(userId)
}
