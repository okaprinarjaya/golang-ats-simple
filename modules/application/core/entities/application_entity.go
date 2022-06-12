package application_core_entities

import (
	"time"

	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
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
		applicationLogs:     nil,
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

	return appl, nil
}

// Business logics

func (appl *ApplicationEntity) MoveToNextStep(nextStepSequence int, stepStatus string) error {
	appl.currentStepSequence = nextStepSequence
	if err := appl.createApplicationLog(stepStatus); err != nil {
		return err
	}
	return nil
}

func (appl *ApplicationEntity) createApplicationLog(stepStatus string) error {
	applLog, err := NewApplicationLogEntity(application_core_dto.ApplicationLogBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        "ApplicationLogId123",
			CreatedAt: time.Now(),
		},
		ApplicationId: appl.Id(),
		JobId:         appl.JobId(),
		StepSequence:  appl.CurrentStepSequence(),
		StepStatus:    stepStatus,
	})

	if err != nil {
		return err
	}

	appl.applicationLogs = append(appl.applicationLogs, applLog)
	return nil
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
