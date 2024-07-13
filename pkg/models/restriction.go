package models

import "time"

type Restriction struct {
	ID        int       `db:"id"`
	RoomID    int       `db:"room_id"`
	Type      string    `db:"type"`
	Effective time.Time `db:"effective"`
	Expire    time.Time `db:"expire"`
}
