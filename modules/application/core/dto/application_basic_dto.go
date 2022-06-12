package application_core_dto

import (
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationBasicDTO struct {
	BaseRecord          core_shared.BaseDTO             `json:"baseRecord"`
	ApplicationLogs     []ApplicationLogBasicDTO        `json:"applicationLogs"`
	ApplicantId         string                          `json:"applicantId"`
	JobId               string                          `json:"jobId"`
	CurrentStepSequence int                             `json:"currentStepSequence"`
	Applicant           application_core_vo.ApplicantVO `json:"applicantDetail"`
	Job                 application_core_vo.JobVO       `json:"jobDetail"`
}
