package application_core_entities

import (
	"time"

	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationLogEntity struct {
	core_shared.BaseEntity
	applicationId                string
	jobId                        string
	hiringStepType               string
	hiringStepTypeCompletedAt    time.Time
	hiringStepSequence           int
	hiringStepStatus             string
	hiringStepStatusClosedAt     time.Time
	hiringStepStatusClosedBy     string
	hiringStepStatusClosedByName string
	userType                     string
}

func NewApplicationLogEntity(applLogDTO application_core_dto.ApplicationLogBasicDTO) ApplicationLogEntity {
	applLog := ApplicationLogEntity{
		applicationId:                applLogDTO.ApplicationId,
		jobId:                        applLogDTO.JobId,
		hiringStepType:               applLogDTO.HiringStepType,
		hiringStepTypeCompletedAt:    applLogDTO.HiringStepTypeCompletedAt,
		hiringStepSequence:           applLogDTO.HiringStepSequence,
		hiringStepStatus:             applLogDTO.HiringStepStatus,
		hiringStepStatusClosedAt:     applLogDTO.HiringStepStatusClosedAt,
		hiringStepStatusClosedBy:     applLogDTO.HiringStepStatusClosedBy,
		hiringStepStatusClosedByName: applLogDTO.HiringStepStatusClosedByName,
		userType:                     applLogDTO.UserType,
	}

	applLog.Base(applLogDTO.BaseRecord)

	return applLog
}

// Business requirements / logics

func (applLog *ApplicationLogEntity) CloseCurrentStatusOfHiringStepType(userBy string, userByName string) {
	applLog.hiringStepTypeCompletedAt = time.Now()
	applLog.hiringStepStatusClosedAt = time.Now()
	applLog.hiringStepStatusClosedBy = userBy
	applLog.hiringStepStatusClosedByName = userByName
	applLog.PersistenceStatus = core_shared.MODIFIED
}

// Data Getters

func (applLog *ApplicationLogEntity) ApplicationId() string {
	return applLog.applicationId
}

func (applLog *ApplicationLogEntity) JobId() string {
	return applLog.jobId
}

func (applLog *ApplicationLogEntity) HiringStepType() string {
	return applLog.hiringStepType
}

func (applLog *ApplicationLogEntity) HiringStepTypeCompletedAt() time.Time {
	return applLog.hiringStepTypeCompletedAt
}

func (applLog *ApplicationLogEntity) HiringStepSequence() int {
	return applLog.hiringStepSequence
}

func (applLog *ApplicationLogEntity) HiringStepStatus() string {
	return applLog.hiringStepStatus
}

func (applLog *ApplicationLogEntity) HiringStepStatusClosedAt() time.Time {
	return applLog.hiringStepStatusClosedAt
}

func (applLog *ApplicationLogEntity) HiringStepStatusClosedBy() string {
	return applLog.hiringStepStatusClosedBy
}

func (applLog *ApplicationLogEntity) HiringStepStatusClosedByName() string {
	return applLog.hiringStepStatusClosedByName
}

func (applLog *ApplicationLogEntity) UserType() string {
	return applLog.userType
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
