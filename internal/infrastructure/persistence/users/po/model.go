package persistence

import "time"

type Users struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}
