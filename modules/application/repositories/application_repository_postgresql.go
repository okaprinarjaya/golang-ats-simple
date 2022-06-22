package application_repositories

import (
	"time"

	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_repositories_datamodels "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/repositories/data-models"
	"gorm.io/gorm"
)

type ApplicationRepositoryPostgreSql struct {
	Db *gorm.DB
}

func NewApplicationRepositoryPostgreSql(db *gorm.DB) *ApplicationRepositoryPostgreSql {
	return &ApplicationRepositoryPostgreSql{Db: db}
}

func (repo *ApplicationRepositoryPostgreSql) Create(applicationEntity application_core_entities.ApplicationEntity) error {
	application := application_repositories_datamodels.Application{
		ID:          "application-id-123abc" + time.Now().String(),
		ApplicantId: applicationEntity.ApplicantId(),
		JobId:       applicationEntity.JobId(),
		CreatedAt:   time.Now(),
	}

	res := repo.Db.Create(&application)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo *ApplicationRepositoryPostgreSql) Update(applicationEntity application_core_entities.ApplicationEntity) error {
	return nil
}
func (repo *ApplicationRepositoryPostgreSql) Delete(applicationEntity application_core_entities.ApplicationEntity) error {
	return nil
}
func (repo *ApplicationRepositoryPostgreSql) FindAllByJobId(jobId string) ([]application_core_entities.ApplicationEntity, error) {
	return nil, nil
}
