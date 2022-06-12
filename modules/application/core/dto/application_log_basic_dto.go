package application_core_dto

import (
	"time"

	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationLogBasicDTO struct {
	BaseRecord    core_shared.BaseDTO `json:"baseRecord"`
	ApplicationId string              `json:"applicationId"`
	JobId         string              `json:"jobId"`
	StepSequence  int                 `json:"stepSequence"`
	StepStatus    string              `json:"stepStatus"`
	CompletedDate time.Time           `json:"completedDate"`
}
