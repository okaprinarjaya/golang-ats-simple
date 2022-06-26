package application_repositories

import application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"

type IApplicationRepository interface {
	Save(applicationEntity application_core_entities.ApplicationEntity) error
	Delete(applicationEntity application_core_entities.ApplicationEntity) error
	FindById(id string) (*application_core_entities.ApplicationEntity, error)
	FindAllByJobId(jobId string) ([]application_core_entities.ApplicationEntity, error)
}
