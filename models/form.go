package models

import "time"

type Form struct {
	ID          uint
	Title       string
	Description string
	Banner      string
	Owner       int64
	Public      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
