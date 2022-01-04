package data

type Event struct {
	Id     int    `json:"id"`
	Name   string `json:"event_name"`
	UserId int    `json:"user_id,omitempty"`
}

type EventSubjects struct {
	EventId  int       `json:"event_id"`
	Subjects []Subject `json:"subjects"`
}
