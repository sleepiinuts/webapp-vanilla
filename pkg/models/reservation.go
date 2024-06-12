package models

import "time"

type Reservation struct {
	ID        int
	UserID    int
	RoomID    int
	Arrival   time.Time
	Departure time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
