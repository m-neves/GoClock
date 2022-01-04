package service

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/repository"
)

type EventServiceInterface interface {
	Create(event *data.Event) error
	AddSubjects(eventSubjects *data.EventSubjects) error
	FindAll(id, userId int) (*data.Event, error)
}

type eventService struct {
	er repository.EventRepositoryInterface
}

func NewEventService(er repository.EventRepositoryInterface) EventServiceInterface {
	return &eventService{er}
}

func (es *eventService) Create(event *data.Event) error {
	return es.er.Create(event.Name, event.UserId)
}

func (es *eventService) AddSubjects(eventSubjects *data.EventSubjects) error {
	return es.er.AddSubjects(eventSubjects)
}

func (es *eventService) FindAll(id, userId int) (*data.Event, error) {
	return es.er.FindAll(id, userId)
}
