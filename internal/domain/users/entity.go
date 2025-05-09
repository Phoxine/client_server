package users

import "time"

type Users struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}
