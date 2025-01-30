package models

import "time"

type Token struct {
	UserID int
	Exp    time.Time
}
