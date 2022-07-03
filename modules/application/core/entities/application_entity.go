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
		job:                       applDTO.Job,
	}

	appl.Base(applDTO.BaseRecord)

	for _, appLogDTO := range applDTO.ApplicationLogs {
		appl.applicationLogs = append(appl.applicationLogs, NewApplicationLogEntity(appLogDTO))
	}

	return &appl, nil
}

// Business requirements / logics

func (appl *ApplicationEntity) MoveFromCVSubmissionToNextStep(
	applLogBaseDTO core_shared.BaseDTO,
	nextHiringStepSequence int,
	hiringStepType string,
	updatedBy string,
	updatedByName string,
	userType string,
) error {
	if appl.isHiringStepLogged(nextHiringStepSequence, constants.APPL_STEP_STATUS_PASSED) {
		return fmt.Errorf("hiring step movement collision (double movement)")
	}

	if appl.currentHiringStepSequence == 1 {
		appl.currentHiringStepSequence = nextHiringStepSequence
		appl.SetUpdatedAt(time.Now())
		appl.SetUpdatedBy(updatedBy)
		appl.SetUpdatedByName(updatedByName)
		appl.PersistenceStatus = core_shared.MODIFIED

		for i := 0; i < len(appl.applicationLogs); i++ {
			applLog := &appl.applicationLogs[i]
			if applLog.hiringStepSequence == 1 &&
				applLog.hiringStepType == constants.HIRING_STEP_TYPE_CV_SUBMISSION &&
				applLog.hiringStepStatus == constants.APPL_STEP_STATUS_IN_PROGRESS {
				applLog.hiringStepTypeCompletedAt = time.Now()
				applLog.hiringStepStatusClosedAt = time.Now()
				applLog.hiringStepStatusClosedBy = updatedBy
				applLog.hiringStepStatusClosedByName = updatedByName
				applLog.PersistenceStatus = core_shared.MODIFIED

				break
			}
		}

		appl.createApplicationLog2(
			applLogBaseDTO,
			1,
			constants.HIRING_STEP_TYPE_CV_SUBMISSION,
			constants.APPL_STEP_STATUS_PASSED,
			userType,
		)

		return nil
	}

	return fmt.Errorf("hiring step sequence is not CV Submission")
}

func (appl *ApplicationEntity) MoveToNextStep(
	applLogBaseRecord core_shared.BaseDTO,
	nextHiringStepSequence int,
	hiringStepType string,
	updatedBy string,
) error {
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

	appl.createApplicationLog(applLogBaseRecord, hiringStepType)
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

func (appl *ApplicationEntity) createApplicationLog(baseRecord core_shared.BaseDTO, hiringStepType string) {
	applLog := NewApplicationLogEntity(application_core_dto.ApplicationLogBasicDTO{
		BaseRecord:         baseRecord,
		ApplicationId:      appl.Id(),
		JobId:              appl.jobId,
		HiringStepType:     hiringStepType,
		HiringStepSequence: appl.currentHiringStepSequence,
		HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
		UserType:           "APPLICANT",
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

func (appl *ApplicationEntity) createApplicationLog2(
	baseRecordDTO core_shared.BaseDTO,
	hiringStepSequence int,
	hiringStepType string,
	hiringStepStatus string,
	userType string,
) {
	applLog := NewApplicationLogEntity(application_core_dto.ApplicationLogBasicDTO{
		BaseRecord:         baseRecordDTO,
		ApplicationId:      appl.Id(),
		JobId:              appl.jobId,
		HiringStepType:     hiringStepType,
		HiringStepSequence: hiringStepSequence,
		HiringStepStatus:   hiringStepStatus,
		UserType:           userType,
	})
	applLog.PersistenceStatus = core_shared.NEW
	appl.applicationLogs = append(appl.applicationLogs, applLog)
}

func (appl *ApplicationEntity) isHiringStepLogged(hiringStepSequence int, hiringStepStatus string) bool {
	logged := false
	for _, applLog := range appl.applicationLogs {
		if applLog.hiringStepSequence == hiringStepSequence && applLog.hiringStepStatus == hiringStepStatus {
			logged = true
			break
		}
	}
	return logged
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

func (appl *ApplicationEntity) IsRejected() bool {
	return appl.isRejected
}

func (appl *ApplicationEntity) IsCancelled() bool {
	return appl.isCancelled
}

func (appl *ApplicationEntity) IsWithdrawed() bool {
	return appl.isWithdrawed
}

func (appl *ApplicationEntity) IsOffered() bool {
	return appl.isOffered
}

func (appl *ApplicationEntity) IsHired() bool {
	return appl.isHired
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
