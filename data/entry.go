package data

import (
	"database/sql"
	"time"
)

type EntryType uint8

const (
	TypeStudy EntryType = iota
	TypePause
	TypePractice
)

type Entry struct {
	Id          uint          `json:"id"`
	Description string        `json:"entry_desc"`
	Type        EntryType     `json:"type"`
	Start       time.Time     `json:"dt_start"`
	End         time.Time     `json:"dt_end"`
	UserId      int           `json:"user_id,omitempty"`
	SubjectId   int           `json:"subject_id,omitempty"`
	EventId     sql.NullInt64 `json:"event_id,omitempty"`
}

func (e *Entry) Diff() time.Duration {
	return e.End.Sub(e.Start)
}

func (e *Entry) IsValid() bool {
	return (!e.Start.IsZero() && !e.End.IsZero()) && e.End.After(e.Start)
}

type Entries []*Entry

func (es *Entries) TotalDiff() time.Duration {
	var total time.Duration

	for _, v := range *es {
		total += v.Diff()
	}

	return total
}
