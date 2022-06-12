package application_core_entities

import (
	"fmt"
	"time"

	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

const (
	APPL_STEP_STATUS_IN_PROGRESS = "IN_PROGRESS"
	APPL_STEP_STATUS_IN_REVIEW   = "IN_REVIEW"
	APPL_STEP_STATUS_PASSED      = "PASSED"
	APPL_STEP_STATUS_REJECTED    = "REJECTED"
	APPL_STEP_STATUS_WITHDRAW    = "WITHDRAW"
	APPL_STEP_STATUS_OFFERED     = "OFFERED"
	APPL_STEP_STATUS_CANCELLED   = "CANCELLED"
	APPL_STEP_STATUS_HIRED       = "HIRED"
)

type ApplicationEntity struct {
	core_shared.BaseEntity
	applicationLogs     []ApplicationLogEntity
	applicantId         string
	jobId               string
	currentStepSequence int
	applicant           application_core_vo.ApplicantVO
	job                 application_core_vo.JobVO
}

func NewApplicationEntity(applDTO application_core_dto.ApplicationBasicDTO) (ApplicationEntity, error) {
	appl := ApplicationEntity{
		applicantId:         applDTO.ApplicantId,
		jobId:               applDTO.JobId,
		currentStepSequence: 1,
		job: application_core_vo.JobVO{
			JobName:           applDTO.Job.JobName,
			JobSAdmtatus:      applDTO.Job.JobSAdmtatus,
			JobDepartmentId:   applDTO.Job.JobDepartmentId,
			JobDepartmentName: applDTO.Job.JobDepartmentName,
			JobCountryId:      applDTO.Job.JobCountryId,
			JobCountryName:    applDTO.Job.JobCountryName,
			JobCityId:         applDTO.Job.JobCityId,
			JobCityName:       applDTO.Job.JobCityName,
		},
	}

	appl.Base(core_shared.BaseDTO{
		Id:        applDTO.BaseRecord.Id,
		CreatedAt: time.Now(),
		CreatedBy: applDTO.BaseRecord.CreatedBy,
	})

	appl.createApplicationLog()

	return appl, nil
}

// Business requirements / logics

func (appl *ApplicationEntity) MoveToNextStep(nextStepSequence int, stepStatus string) error {
	for _, v := range appl.applicationLogs {
		if v.stepSequence == nextStepSequence {
			return fmt.Errorf("duplicated step sequence")
		}
	}

	appl.currentStepSequence = nextStepSequence
	appl.PersistenceStatus = core_shared.MODIFIED

	appl.createApplicationLog()
	return nil
}

func (appl *ApplicationEntity) UpdateStepStatus(targettedStepSequence int, newStepStatus string, updatedBy string) error {
	found := false
	for i := 0; i < len(appl.applicationLogs); i++ {
		if appl.applicationLogs[i].stepSequence == targettedStepSequence {
			found = true
			var appLog *ApplicationLogEntity = &appl.applicationLogs[i]

			if newStepStatus == APPL_STEP_STATUS_REJECTED || newStepStatus == APPL_STEP_STATUS_WITHDRAW || newStepStatus == APPL_STEP_STATUS_CANCELLED || newStepStatus == APPL_STEP_STATUS_HIRED {
				appLog.completedDate = time.Now()
			}
			appLog.stepStatus = newStepStatus
			appLog.BaseEntity.SetUpdatedAt(time.Now())
			appLog.BaseEntity.SetUpdatedBy(updatedBy)
			appLog.PersistenceStatus = core_shared.MODIFIED
			break
		}
	}

	if !found {
		return fmt.Errorf("step sequence not found")
	}

	return nil
}

func (appl *ApplicationEntity) createApplicationLog() {
	applLog := NewApplicationLogEntity(application_core_dto.ApplicationLogBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        "ApplicationLogId123",
			CreatedAt: time.Now(),
			CreatedBy: appl.CreatedBy(),
		},
		ApplicationId: appl.Id(),
		JobId:         appl.jobId,
		StepSequence:  appl.currentStepSequence,
		StepStatus:    APPL_STEP_STATUS_IN_PROGRESS,
	})
	applLog.PersistenceStatus = core_shared.NEW

	if len(appl.applicationLogs) > 1 {
		prevStepSequence := appl.currentStepSequence - 1
		for i := 0; i < len(appl.applicationLogs); i++ {
			if appl.applicationLogs[i].stepSequence == prevStepSequence {
				var applLog *ApplicationLogEntity = &appl.applicationLogs[i]
				applLog.stepStatus = APPL_STEP_STATUS_PASSED
				applLog.BaseEntity.SetUpdatedAt(time.Now())
				applLog.BaseEntity.SetUpdatedBy(appl.CreatedBy())
				applLog.PersistenceStatus = core_shared.MODIFIED
				break
			}
		}
	}

	appl.applicationLogs = append(appl.applicationLogs, applLog)
}

// Data Getters

func (appl *ApplicationEntity) ApplicationLogs() []ApplicationLogEntity {
	return appl.applicationLogs
}

func (appl *ApplicationEntity) ApplicantId() string {
	return appl.applicantId
}

func (appl *ApplicationEntity) JobId() string {
	return appl.jobId
}

func (appl *ApplicationEntity) CurrentStepSequence() int {
	return appl.currentStepSequence
}

func (appl *ApplicationEntity) Applicant() application_core_vo.ApplicantVO {
	return appl.applicant
}

func (appl *ApplicationEntity) Job() application_core_vo.JobVO {
	return appl.job
}

func (appl *ApplicationEntity) Id() string {
	return appl.BaseEntity.Id()
}

func (appl *ApplicationEntity) CreatedAt() time.Time {
	return appl.BaseEntity.CreatedAt()
}

func (appl *ApplicationEntity) UpdatedAt() time.Time {
	return appl.BaseEntity.UpdatedAt()
}

func (appl *ApplicationEntity) DeletedAt() time.Time {
	return appl.BaseEntity.DeletedAt()
}

func (appl *ApplicationEntity) CreatedBy() string {
	return appl.BaseEntity.CreatedBy()
}

func (appl *ApplicationEntity) UpdatedBy() string {
	return appl.BaseEntity.UpdatedBy()
}

func (appl *ApplicationEntity) DeletedBy() string {
	return appl.BaseEntity.DeletedBy()
}
