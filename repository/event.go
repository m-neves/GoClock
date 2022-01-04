package repository

import (
	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/database"
)

type EventRepositoryInterface interface {
	Create(name string, userId int) error
	AddSubjects(eventSubjects *data.EventSubjects) error
	FindAll(id, userId int) (*data.Event, error)
}

type eventRepository struct {
}

func NewEventRepository() *eventRepository {
	return &eventRepository{}
}

func (er *eventRepository) Create(name string, userId int) error {
	_, err := database.GetDb().Exec(`
		INSERT INTO goclock.tbl_event
			(event_name, user_id)
		VALUES
			($1, $2);
	`, name, userId)

	return err
}

func (er *eventRepository) AddSubjects(eventSubjects *data.EventSubjects) error {

	for _, v := range eventSubjects.Subjects {

		_, err := database.GetDb().Exec(`
			INSERT INTO goclock.tbl_event_subjects
				(event_id, subject_id)
			VALUES
				($1, $2);
		`, eventSubjects.EventId, v.Id)

		if err != nil {
			return err
		}

	}

	return nil
}

func (er *eventRepository) FindAll(id, userId int) (*data.Event, error) {
	row := database.GetDb().QueryRow(`
		SELECT e.id, e.name, e.user_id
		FROM tbl_event e
		WHERE e.id = $1
		AND e.user_id = $2
	`, id, userId)

	event := &data.Event{}

	err := row.Scan(&event.Id, &event.UserId)

	if err != nil {
		return nil, err
	}

	return event, nil
}
