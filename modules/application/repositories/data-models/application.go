package application_repositories_datamodels

import "time"

type Application struct {
	ID          string `gorm:"primaryKey"`
	ApplicantId string
	JobId       string
	CreatedAt   time.Time
}
