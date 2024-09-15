package entity

import "time"

type Task struct {
	Id          uint16
	Description string
	CreatedAt   time.Time
	IsDone      bool
}
