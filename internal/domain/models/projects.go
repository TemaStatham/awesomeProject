package models

import "time"

type Projects struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	CreatedAt time.Duration `json:"created_at"`
}
