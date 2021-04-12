package models

type Meetup struct {
	ID          string `json:"id" pg:",pk,unique,notnull"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        *User  `pg:"fk:user_id"`
	UserID      string `json:"user_id"`
}
