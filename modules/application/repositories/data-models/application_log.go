package application_repositories_datamodels

import (
	"database/sql"
	"time"
)

type ApplicationLog struct {
	ID                           string         `gorm:"primaryKey"`
	ApplicationId                string         `gorm:"column:application_id"`
	JobId                        string         `gorm:"column:job_id"`
	HiringStepType               string         `gorm:"column:hiring_step_type"`
	HiringStepSequence           int            `gorm:"column:hiring_step_sequence"`
	HiringStepStatus             string         `gorm:"column:hiring_step_status"`
	HiringStepStatusClosedAt     sql.NullTime   `gorm:"column:hiring_step_status_closed_at"`
	HiringStepStatusClosedBy     sql.NullString `gorm:"column:hiring_step_status_closed_by"`
	HiringStepStatusClosedByName sql.NullString `gorm:"column:hiring_step_status_closed_by_name"`
	UserType                     string         `gorm:"column:user_type"`
	CreatedAt                    time.Time      `gorm:"column:created_at"`
	CreatedBy                    string         `gorm:"column:created_by"`
	CreatedByName                string         `gorm:"column:created_by_name"`
}
