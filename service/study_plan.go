package service

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/repository"
)

type StudyPlanServiceInterface interface {
	Create(studyPlan *data.StudyPlan) error
	AddCourse(studyPlanCourse *data.StudyPlanCourse) error
	FindById(id, user_id uint) (*data.StudyPlanCourses, error)
}

type studyPlanService struct {
	r repository.StudyPlanRepositoryInterface
}

func NewStudyPlanService(r repository.StudyPlanRepositoryInterface) StudyPlanServiceInterface {
	return &studyPlanService{r}
}

func (sps *studyPlanService) Create(studyPlan *data.StudyPlan) error {
	return sps.r.Create(studyPlan)
}

func (sps *studyPlanService) AddCourse(studyPlanCourse *data.StudyPlanCourse) error {
	return sps.r.AddCourse(studyPlanCourse)
}

func (sps *studyPlanService) FindById(id, user_id uint) (*data.StudyPlanCourses, error) {
	return sps.r.FindById(id, user_id)
}
