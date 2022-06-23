package core_shared

import "time"

type BaseDTO struct {
	Id            string    `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt"`
	CreatedBy     string    `json:"createdBy"`
	CreatedByName string    `json:"createdByName"`
	UpdatedBy     string    `json:"updatedBy"`
	UpdatedByName string    `json:"updatedByName"`
	DeletedBy     string    `json:"deletedBy"`
	DeletedByName string    `json:"deletedByName"`
}
