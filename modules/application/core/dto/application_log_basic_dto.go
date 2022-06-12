package application_core_dto

import (
	"time"

	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationLogBasicDTO struct {
	BaseRecord                core_shared.BaseDTO `json:"baseRecord"`
	ApplicationId             string              `json:"applicationId"`
	JobId                     string              `json:"jobId"`
	HiringStepType            string              `json:"hiringStepType"`
	HiringSteptypeCompletedAt time.Time           `json:"hiringStepTypeCompletedAt"`
	StepSequence              int                 `json:"stepSequence"`
	StepStatus                string              `json:"stepStatus"`
}
