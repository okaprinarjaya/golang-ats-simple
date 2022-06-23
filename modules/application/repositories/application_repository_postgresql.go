package application_repositories

import (
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_repositories_datamodels "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/repositories/data-models"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
	"gorm.io/gorm"
)

type ApplicationRepositoryPostgreSql struct {
	Db *gorm.DB
}

func NewApplicationRepositoryPostgreSql(db *gorm.DB) *ApplicationRepositoryPostgreSql {
	return &ApplicationRepositoryPostgreSql{Db: db}
}

func (repo *ApplicationRepositoryPostgreSql) Save(applicationEntity application_core_entities.ApplicationEntity) error {
	application := transformBusinessEntityToDataModel(applicationEntity)

	if applicationEntity.PersistenceStatus == core_shared.NEW {
		res := repo.Db.Create(&application)

		if res.Error != nil {
			return res.Error
		}
	}

	if applicationEntity.PersistenceStatus == core_shared.MODIFIED {
		res := repo.Db.Updates(&application)

		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (repo *ApplicationRepositoryPostgreSql) Delete(applicationEntity application_core_entities.ApplicationEntity) error {
	return nil
}

func (repo *ApplicationRepositoryPostgreSql) FindAllByJobId(jobId string) ([]application_core_entities.ApplicationEntity, error) {
	return nil, nil
}

func transformBusinessEntityToDataModel(applicationEntity application_core_entities.ApplicationEntity) application_repositories_datamodels.Application {
	application := application_repositories_datamodels.Application{
		ID:                        applicationEntity.Id(),
		ApplicantId:               applicationEntity.ApplicantId(),
		JobId:                     applicationEntity.JobId(),
		CurrentHiringStepSequence: applicationEntity.CurrentHiringStepSequence(),
		IsRejected:                applicationEntity.IsRejected(),
		IsCancelled:               applicationEntity.IsCancelled(),
		IsWithdrawed:              applicationEntity.IsWithdrawed(),
		IsOffered:                 applicationEntity.IsOffered(),
		IsHired:                   applicationEntity.IsHired(),

		ApplicantCompleteName:           applicationEntity.Applicant().ApplicantCompleteName,
		ApplicantGender:                 applicationEntity.Applicant().ApplicantGender,
		ApplicantDateOfBirth:            applicationEntity.Applicant().ApplicantDateOfBirth,
		ApplicantAddress:                applicationEntity.Applicant().ApplicantAddress,
		ApplicantProfilePhoto:           applicationEntity.Applicant().ApplicantProfilePhoto,
		ApplicantProfileSummary:         applicationEntity.Applicant().ApplicantProfileSummary,
		ApplicantNationality:            applicationEntity.Applicant().ApplicantNationality,
		ApplicantCountryId:              applicationEntity.Applicant().ApplicantCountryId,
		ApplicantCountryName:            applicationEntity.Applicant().ApplicantCountryName,
		ApplicantCityId:                 applicationEntity.Applicant().ApplicantCityId,
		ApplicantCityName:               applicationEntity.Applicant().ApplicantCityName,
		ApplicantIsLookingOppty:         applicationEntity.Applicant().ApplicantIsLookingOppty,
		ApplicantEducationLast:          applicationEntity.Applicant().ApplicantEducationLast,
		ApplicantTotalYearsXp:           applicationEntity.Applicant().ApplicantTotalYearsXp,
		ApplicantJobLevelLast:           applicationEntity.Applicant().ApplicantJobLevelLast,
		ApplicantWillingWorkOverseas:    applicationEntity.Applicant().ApplicantWillingWorkOverseas,
		ApplicantExpectedSalary:         applicationEntity.Applicant().ApplicantExpectedSalary,
		ApplicantExpectedSalaryCurrency: applicationEntity.Applicant().ApplicantExpectedSalaryCurrency,

		JobName:           applicationEntity.Job().JobName,
		JobAdmStatus:      applicationEntity.Job().JobAdmStatus,
		JobDepartmentId:   applicationEntity.Job().JobDepartmentId,
		JobDepartmentName: applicationEntity.Job().JobDepartmentName,
		JobCountryId:      applicationEntity.Job().JobCountryId,
		JobCountryName:    applicationEntity.Job().JobCountryName,
		JobCityId:         applicationEntity.Job().JobCityId,
		JobCityName:       applicationEntity.Job().JobCityName,

		CreatedAt:     applicationEntity.CreatedAt(),
		UpdatedAt:     applicationEntity.UpdatedAt(),
		DeletedAt:     applicationEntity.DeletedAt(),
		CreatedBy:     applicationEntity.CreatedBy(),
		CreatedByName: applicationEntity.CreatedByName(),
		UpdatedBy:     applicationEntity.UpdatedBy(),
		UpdatedByName: applicationEntity.UpdatedByName(),
		DeletedBy:     applicationEntity.DeletedBy(),
		DeletedByName: applicationEntity.DeletedByName(),
	}

	return application
}
