package repository

import (
	"time"

	"github.com/m-neves/goclock/data"
	"github.com/m-neves/goclock/database"
)

type EntryRepositoryInterface interface {
	Create(entry *data.Entry) error
	FindById(id, userId int) (*data.Entry, error)
	SumByPeriod(start, end time.Time, userId int) (*data.Entries, error)
}

type entryRepository struct {
}

func NewEntryRepository() EntryRepositoryInterface {
	return &entryRepository{}
}

func (er *entryRepository) Create(entry *data.Entry) error {
	_, err := database.GetDb().Exec(`
		INSERT INTO goclock.tbl_entry
			(entry_desc, type, user_id, subject_id, dt_start, dt_end)
		VALUES
			($1, $2, $3, $4, $5, $6);
	`, entry.Description, entry.Type, entry.UserId, entry.SubjectId, entry.Start, entry.End,
	)

	return err
}

func (er *entryRepository) FindById(id, userId int) (*data.Entry, error) {
	row := database.GetDb().QueryRow(`
		SELECT e.id, e.dt_start, e.dt_end, e.user_id, e.subject_id
		FROM tbl_entry e
		WHERE e.id = $1
		AND e.user_id = $2
	`, id, userId)

	entry := &data.Entry{}

	err := row.Scan(&entry.Id, &entry.Start, &entry.End, &entry.UserId, &entry.SubjectId)

	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (er *entryRepository) SumByPeriod(start, end time.Time, userId int) (*data.Entries, error) {

	rows, err := database.GetDb().Query(`
		SELECT e.id, e.subject_id, e.event_id, e.user_id, e.dt_start, e.dt_end
		FROM tbl_entry e
		WHERE e.user_id = $1
		AND e.dt_start >= $2
		AND e.dt_end <= $3
	`, userId, start, end)

	if err != nil {
		return nil, err
	}

	entries := &data.Entries{}

	for rows.Next() {
		entry := &data.Entry{}

		err := rows.Scan(
			&entry.Id,
			&entry.SubjectId,
			&entry.EventId,
			&entry.UserId,
			&entry.Start,
			&entry.End,
		)

		if err != nil {
			return nil, err
		}

		*entries = append(*entries, entry)
	}

	return entries, nil
}
