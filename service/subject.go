package service

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/repository"
)

type SubjectServiceInterface interface {
	FindById(id, userId int) (*data.Subject, error)
	FindAll(userId int) ([]*data.Subject, error)
	Create(name string, userId int) error
}

type subjectService struct {
	sr repository.SubjectRepositoryInterface
}

func NewSubjectService(sr repository.SubjectRepositoryInterface) *subjectService {
	return &subjectService{sr}
}

func (ss *subjectService) FindById(id, userId int) (*data.Subject, error) {
	return ss.sr.FindById(id, userId)
}

func (ss *subjectService) FindAll(userId int) ([]*data.Subject, error) {
	return ss.sr.FindAll(userId)
}

func (ss *subjectService) Create(name string, userId int) error {
	return ss.sr.Create(name, userId)
}
