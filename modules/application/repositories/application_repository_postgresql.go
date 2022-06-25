package application_repositories

import (
	"database/sql"

	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_repositories_datamodels "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/repositories/data-models"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
	"gitlab.com/okaprinarjaya.wartek/ats-simple/utils"
	"gorm.io/gorm"
)

type ApplicationRepositoryPostgreSql struct {
	Db *gorm.DB
}

func NewApplicationRepositoryPostgreSql(db *gorm.DB) *ApplicationRepositoryPostgreSql {
	return &ApplicationRepositoryPostgreSql{Db: db}
}

func (repo *ApplicationRepositoryPostgreSql) Save(applicationEntity application_core_entities.ApplicationEntity) error {
	application, applicationLogList := transformBusinessEntityToDataModel(applicationEntity)

	return repo.Db.Transaction(func(tx *gorm.DB) error {
		if applicationEntity.PersistenceStatus == core_shared.NEW {
			res := tx.Create(&application)

			if res.Error != nil {
				return res.Error
			}
		}

		if applicationEntity.PersistenceStatus == core_shared.MODIFIED {
			res := tx.Updates(&application)

			if res.Error != nil {
				return res.Error
			}
		}

		if len(applicationLogList) > 0 {
			res := tx.Create(&applicationLogList[0])

			if res.Error != nil {
				return res.Error
			}
		}

		return nil
	})
}

func (repo *ApplicationRepositoryPostgreSql) Delete(applicationEntity application_core_entities.ApplicationEntity) error {
	return nil
}

func (repo *ApplicationRepositoryPostgreSql) FindAllByJobId(jobId string) ([]application_core_entities.ApplicationEntity, error) {
	return nil, nil
}

func transformBusinessEntityToDataModel(applicationEntity application_core_entities.ApplicationEntity) (
	application_repositories_datamodels.Application,
	[]application_repositories_datamodels.ApplicationLog,
) {
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

		CreatedAt: applicationEntity.CreatedAt(),
		UpdatedAt: sql.NullTime{
			Time:  applicationEntity.UpdatedAt(),
			Valid: utils.DateValid(applicationEntity.UpdatedAt()),
		},
		DeletedAt: sql.NullTime{
			Time:  applicationEntity.DeletedAt(),
			Valid: utils.DateValid(applicationEntity.DeletedAt()),
		},
		CreatedBy:     applicationEntity.CreatedBy(),
		CreatedByName: applicationEntity.CreatedByName(),
		UpdatedBy: sql.NullString{
			String: applicationEntity.UpdatedBy(),
			Valid:  utils.StringValid(applicationEntity.UpdatedBy()),
		},
		UpdatedByName: sql.NullString{
			String: applicationEntity.UpdatedByName(),
			Valid:  utils.StringValid(applicationEntity.UpdatedByName()),
		},
		DeletedBy: sql.NullString{
			String: applicationEntity.DeletedBy(),
			Valid:  utils.StringValid(applicationEntity.DeletedBy()),
		},
		DeletedByName: sql.NullString{
			String: applicationEntity.DeletedByName(),
			Valid:  utils.StringValid(applicationEntity.DeletedByName()),
		},
	}

	var applicationLogList []application_repositories_datamodels.ApplicationLog

	for _, log := range applicationEntity.ApplicationLogs() {
		if log.PersistenceStatus == core_shared.NEW || log.PersistenceStatus == core_shared.MODIFIED {
			applicationLog := application_repositories_datamodels.ApplicationLog{
				ID:             log.Id(),
				ApplicationId:  log.ApplicationId(),
				JobId:          log.JobId(),
				HiringStepType: log.HiringStepType(),
				HiringStepTypeCompletedAt: sql.NullTime{
					Time:  log.HiringStepTypeCompletedAt(),
					Valid: utils.DateValid(log.HiringStepTypeCompletedAt()),
				},
				HiringStepSequence: log.HiringStepSequence(),
				HiringStepStatus:   log.HiringStepStatus(),
				HiringStepStatusClosedAt: sql.NullTime{
					Time:  log.HiringStepStatusClosedAt(),
					Valid: utils.DateValid(log.HiringStepStatusClosedAt()),
				},
				HiringStepStatusClosedBy: sql.NullString{
					String: log.HiringStepStatusClosedBy(),
					Valid:  utils.StringValid(log.HiringStepStatusClosedBy()),
				},
				HiringStepStatusClosedByName: sql.NullString{
					String: log.HiringStepStatusClosedByName(),
					Valid:  utils.StringValid(log.HiringStepStatusClosedByName()),
				},
				UserType:      log.UserType(),
				CreatedAt:     log.CreatedAt(),
				CreatedBy:     log.CreatedBy(),
				CreatedByName: log.CreatedByName(),
			}
			applicationLogList = append(applicationLogList, applicationLog)
		}
	}

	return application, applicationLogList
}
