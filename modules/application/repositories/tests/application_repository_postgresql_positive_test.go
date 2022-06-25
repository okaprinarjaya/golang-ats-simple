package application_repositories_tests

import (
	"path"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	infrastructure_database "gitlab.com/okaprinarjaya.wartek/ats-simple/infrastructure/database"

	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_repositories "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/repositories"
	application_repositories_datamodels "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/repositories/data-models"
	application_repositories_tests_fixtures "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/repositories/tests/fixtures"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
	"gorm.io/gorm"
)

type ApplicationRepositoryPostgreSqlPositiveTestSuite struct {
	suite.Suite
	Db *gorm.DB
}

func TestApplicationRepositoryPostgreSqlPositiveTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationRepositoryPostgreSqlPositiveTestSuite))
}

func (s *ApplicationRepositoryPostgreSqlPositiveTestSuite) SetupSuite() {
	err := infrastructure_database.ManageDatabase()
	require.NoError(s.T(), err)

	_, filename, _, _ := runtime.Caller(0)
	migration, errm := infrastructure_database.Migration(path.Join(path.Dir(filename), "../../migrations"))
	require.NoError(s.T(), errm)

	errmup := migration.Up()
	require.NoError(s.T(), errmup)

	s.Db = infrastructure_database.PgConnectionTesting("")
}

func (s *ApplicationRepositoryPostgreSqlPositiveTestSuite) TearDownSuite() {
	sql, _ := s.Db.DB()
	sql.Close()
}

func (s *ApplicationRepositoryPostgreSqlPositiveTestSuite) TestCreateNewApplicationData() {
	var repo application_repositories.IApplicationRepository = application_repositories.NewApplicationRepositoryPostgreSql(s.Db)

	dataIdentifier := application_repositories_tests_fixtures.ApplicationDataSample_DTO{
		ApplicationID:               "application-id-001",
		ApplicationCreatedAt:        time.Date(2022, time.May, 1, 10, 0, 0, 0, time.UTC),
		ApplicationCreatedBy:        "APPLICANT-001",
		ApplicationCreatedByName:    "Brewok The Applicant",
		ApplicantID:                 "APPLICANT-001",
		JobID:                       "JOB-001",
		ApplicationLogID:            "appl-log-001",
		ApplicationLogCreatedAt:     time.Date(2022, time.May, 1, 10, 0, 0, 0, time.UTC),
		ApplicationLogCreatedBy:     "APPLICANT-001",
		ApplicationLogCreatedByName: "Brewok The Applicant",
	}

	sampleData := application_repositories_tests_fixtures.ApplicationDataSample1_DTO(dataIdentifier)
	appl, _ := application_core_entities.NewApplicationEntity(sampleData)
	appl.PersistenceStatus = core_shared.NEW

	for i := 0; i < len(appl.ApplicationLogs()); i++ {
		appl.ApplicationLogs()[i].PersistenceStatus = core_shared.NEW
	}

	err := repo.Save(*appl)

	// Make sure new record are successfully added
	s.Nil(err)

	// Make sure newly added application data record successfully retrieved and contains correct information
	var applDataModel application_repositories_datamodels.Application
	result := s.Db.Find(&applDataModel, "id = ?", dataIdentifier.ApplicationID)

	s.NotZero(result.RowsAffected)
	s.Equal(dataIdentifier.ApplicationID, applDataModel.ID)
	s.Equal(dataIdentifier.ApplicantID, applDataModel.ApplicantId)
	s.Equal(dataIdentifier.JobID, applDataModel.JobId)

	// Make sure newly added application_log data record successfully retrieved and contains correct information
	var applLogDataModel application_repositories_datamodels.ApplicationLog
	resultApplLog := s.Db.Where("application_id = ?", dataIdentifier.ApplicationID).Find(&applLogDataModel)

	s.NotZero(resultApplLog.RowsAffected)
	s.Equal(1, int(result.RowsAffected))
	s.Equal(dataIdentifier.ApplicationLogID, applLogDataModel.ID)
	s.Equal(dataIdentifier.ApplicationID, applLogDataModel.ApplicationId)
}

func (s *ApplicationRepositoryPostgreSqlPositiveTestSuite) TestFindApplicationData() {
	application_repositories_tests_fixtures.SeedApplicationsDataSample(s.Db)
	s.Equal(true, true)
}
