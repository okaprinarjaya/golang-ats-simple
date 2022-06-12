package application_core_entities

import (
	"time"

	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationLogEntity struct {
	core_shared.BaseEntity
	applicationId string
	jobId         string
	stepSequence  int
	stepStatus    string
	completedDate time.Time
}

func NewApplicationLogEntity(applLogDTO application_core_dto.ApplicationLogBasicDTO) (ApplicationLogEntity, error) {
	applLog := ApplicationLogEntity{
		applicationId: applLogDTO.ApplicationId,
		jobId:         applLogDTO.JobId,
		stepSequence:  applLogDTO.StepSequence,
		stepStatus:    applLogDTO.StepStatus,
		completedDate: applLogDTO.CompletedDate,
	}

	applLog.Base(core_shared.BaseDTO{
		Id:        applLogDTO.BaseRecord.Id,
		CreatedAt: time.Now(),
		CreatedBy: applLogDTO.BaseRecord.CreatedBy,
	})

	return applLog, nil
}

func (applLog *ApplicationLogEntity) ApplicationId() string {
	return applLog.applicationId
}

func (applLog *ApplicationLogEntity) JobId() string {
	return applLog.jobId
}

func (applLog *ApplicationLogEntity) StepSequence() int {
	return applLog.stepSequence
}

func (applLog *ApplicationLogEntity) StepStatus() string {
	return applLog.stepStatus
}

func (applLog *ApplicationLogEntity) CompletedDate() time.Time {
	return applLog.completedDate
}

func (applLog *ApplicationLogEntity) Id() string {
	return applLog.BaseEntity.Id()
}

func (applLog *ApplicationLogEntity) CreatedAt() time.Time {
	return applLog.BaseEntity.CreatedAt()
}

func (applLog *ApplicationLogEntity) UpdatedAt() time.Time {
	return applLog.BaseEntity.UpdatedAt()
}

func (applLog *ApplicationLogEntity) DeletedAt() time.Time {
	return applLog.BaseEntity.DeletedAt()
}

func (applLog *ApplicationLogEntity) CreatedBy() string {
	return applLog.BaseEntity.CreatedBy()
}

func (applLog *ApplicationLogEntity) UpdatedBy() string {
	return applLog.BaseEntity.UpdatedBy()
}

func (applLog *ApplicationLogEntity) DeletedBy() string {
	return applLog.BaseEntity.DeletedBy()
}
