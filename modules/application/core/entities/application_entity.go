package application_core_entities

import (
	"fmt"
	"time"

	constants "gitlab.com/okaprinarjaya.wartek/ats-simple/modules"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationEntity struct {
	core_shared.BaseEntity
	applicationLogs           []ApplicationLogEntity
	applicantId               string
	jobId                     string
	currentHiringStepSequence int
	isRejected                bool
	isCancelled               bool
	isWithdrawed              bool
	isOffered                 bool
	isHired                   bool
	applicant                 application_core_vo.ApplicantVO
	job                       application_core_vo.JobVO
}

func NewApplicationEntity(applDTO application_core_dto.ApplicationBasicDTO) (*ApplicationEntity, error) {
	if len(applDTO.ApplicationLogs) == 0 {
		return nil, fmt.Errorf("please provide application log movement record")
	}

	appl := ApplicationEntity{
		applicantId:               applDTO.ApplicantId,
		jobId:                     applDTO.JobId,
		currentHiringStepSequence: applDTO.CurrentHiringStepSequence,
		isRejected:                applDTO.IsRejected,
		isCancelled:               applDTO.IsCancelled,
		isWithdrawed:              applDTO.IsWithdrawed,
		isOffered:                 applDTO.IsOffered,
		isHired:                   applDTO.IsHired,
		applicant:                 applDTO.Applicant,
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

	appl.Base(applDTO.BaseRecord)

	for _, appLogDTO := range applDTO.ApplicationLogs {
		appl.applicationLogs = append(appl.applicationLogs, NewApplicationLogEntity(appLogDTO))
	}

	return &appl, nil
}

// Business requirements / logics

func (appl *ApplicationEntity) MoveToNextStep(nextHiringStepSequence int, hiringStepType string, updatedBy string) error {
	for _, v := range appl.applicationLogs {
		if v.hiringStepSequence == nextHiringStepSequence {
			return fmt.Errorf("duplicated step sequence")
		}
	}

	if appl.isRejected || appl.isCancelled || appl.isWithdrawed || appl.isOffered || appl.isHired {
		return fmt.Errorf("application criteria does meet the requirements to be moved to next step")
	}

	appl.currentHiringStepSequence = nextHiringStepSequence
	appl.SetUpdatedAt(time.Now())
	appl.SetUpdatedBy(updatedBy)
	appl.PersistenceStatus = core_shared.MODIFIED

	appl.createApplicationLog(hiringStepType)
	return nil
}

func (appl *ApplicationEntity) UpdateHiringStepStatus(targettedHiringStepSequence int, newHiringStepStatus string, updatedBy string) error {
	found := false
	for i := 0; i < len(appl.applicationLogs); i++ {
		if appl.applicationLogs[i].hiringStepSequence == targettedHiringStepSequence {
			found = true
			var appLog *ApplicationLogEntity = &appl.applicationLogs[i]

			appLog.hiringStepStatus = newHiringStepStatus
			appLog.BaseEntity.SetUpdatedAt(time.Now())
			appLog.BaseEntity.SetUpdatedBy(updatedBy)
			appLog.PersistenceStatus = core_shared.MODIFIED

			break
		}
	}

	if !found {
		return fmt.Errorf("step sequence not found")
	}

	appl.updateApplicationDecisionFlagging(newHiringStepStatus)

	return nil
}

func (appl *ApplicationEntity) updateApplicationDecisionFlagging(newStepStatus string) {
	switch newStepStatus {
	case constants.APPL_STEP_STATUS_REJECTED:
		appl.isRejected = true
	case constants.APPL_STEP_STATUS_CANCELLED:
		appl.isCancelled = true
	case constants.APPL_STEP_STATUS_WITHDRAW:
		appl.isWithdrawed = true
	case constants.APPL_STEP_STATUS_OFFERED:
		appl.isOffered = true
	case constants.APPL_STEP_STATUS_HIRED:
		appl.isHired = true

	}
}

func (appl *ApplicationEntity) createApplicationLog(hiringStepType string) {
	applLog := NewApplicationLogEntity(application_core_dto.ApplicationLogBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        "ApplicationLogId123",
			CreatedAt: time.Now(),
			CreatedBy: appl.CreatedBy(),
		},
		ApplicationId:      appl.Id(),
		JobId:              appl.jobId,
		HiringStepType:     hiringStepType,
		HiringStepSequence: appl.currentHiringStepSequence,
		HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
	})
	applLog.PersistenceStatus = core_shared.NEW

	nextStepIs2ndStep := (appl.currentHiringStepSequence - 1) == 1
	if len(appl.applicationLogs) == 1 && nextStepIs2ndStep {
		var applLog *ApplicationLogEntity = &appl.applicationLogs[0]
		applLog.hiringStepStatus = constants.APPL_STEP_STATUS_PASSED
		applLog.PersistenceStatus = core_shared.MODIFIED
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

func (appl *ApplicationEntity) CurrentHiringStepSequence() int {
	return appl.currentHiringStepSequence
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
