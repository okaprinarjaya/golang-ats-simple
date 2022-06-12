package core_shared

import "time"

type BaseDTO struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
