package domain

import "time"

type Cake struct {
	Id          int
	Title       string
	Description string
	Rating      float64
	Image       string

	UpdatedAt time.Time
}
