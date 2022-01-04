package data

type Subject struct {
	Id     int    `json:"id"`
	Name   string `json:"subject_name"`
	UserId int    `json:"user_id,omitempty"`
}
