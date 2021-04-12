package models

type Meetup struct {
	ID          string `json:"id" pg:",pk"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `pg:"fk:user_id"`
}
