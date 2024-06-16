package models

import "time"

type Reservation struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	RoomID    int       `db:"room_id"`
	Arrival   time.Time `db:"arrival"`
	Departure time.Time `db:"departure"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
