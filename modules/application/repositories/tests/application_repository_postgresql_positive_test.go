package application_repositories_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

type ApplicationRepositoryPostgreSqlTestSuite struct {
	suite.Suite
	Db *gorm.DB
}

func (s *ApplicationRepositoryPostgreSqlTestSuite) SetupSuite() {
	err := infrastructure_database.CreateDatabaseTesting()
	require.NoError(s.T(), err)

	// _, filename, _, _ := runtime.Caller(0)
	// migration, errm := infrastructure_database.Migration(path.Join(path.Dir(filename), "../../migrations"))
	// require.NoError(s.T(), errm)

	// errmup := migration.Up()
	// require.NoError(s.T(), errmup)

	s.Db = infrastructure_database.PgConnectionTesting("")
}

// func (s *ApplicationRepositoryPostgreSqlTestSuite) TearDownSuite() {
// 	infrastructure_database.DropDatabaseTesting()
// }

func TestApplicationRepositoryPostgreSqlPositiveTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationRepositoryPostgreSqlTestSuite))
}

func (s *ApplicationRepositoryPostgreSqlTestSuite) TestSomething() {
	var repo application_repositories.IApplicationRepository = application_repositories.NewApplicationRepositoryPostgreSql(s.Db)

	sampleData := application_repositories_tests_fixtures.DataSample1()
	appl, _ := application_core_entities.NewApplicationEntity(sampleData)
	appl.PersistenceStatus = core_shared.NEW

	for i := 0; i < len(appl.ApplicationLogs()); i++ {
		appl.ApplicationLogs()[i].PersistenceStatus = core_shared.NEW
	}

	err := repo.Save(*appl)

	assert.Nil(s.T(), err)

	var applDataModel application_repositories_datamodels.Application
	result := s.Db.Find(&applDataModel, "id = ?", sampleData.BaseRecord.Id)

	assert.NotZero(s.T(), result.RowsAffected)
	assert.Equal(s.T(), sampleData.BaseRecord.Id, applDataModel.ID)
}
