package application_core_dto

import (
	"time"

	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationLogBasicDTO struct {
	BaseRecord                   core_shared.BaseDTO `json:"baseRecord"`
	ApplicationId                string              `json:"applicationId"`
	JobId                        string              `json:"jobId"`
	HiringStepType               string              `json:"hiringStepType"`
	HiringStepSequence           int                 `json:"hiringStepSequence"`
	HiringStepStatus             string              `json:"hiringStepStatus"`
	HiringStepStatusClosedAt     time.Time           `json:"hiringStepStatusClosedAt"`
	HiringStepStatusClosedBy     string              `json:"hiringStepStatusClosedBy"`
	HiringStepStatusClosedByName string              `json:"hiringStepStatusClosedByName"`
	UserType                     string              `json:"userType"`
}
