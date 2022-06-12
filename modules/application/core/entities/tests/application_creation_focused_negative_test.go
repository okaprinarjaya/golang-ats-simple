package application_core_entities_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_core_entities_tests_fixtures "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities/tests/fixtures"
)

func TestCreateApplication_InitialCvSubmissionInProgress_NegativeCase(t *testing.T) {
	applData := application_core_entities_tests_fixtures.CreateApplicationData()
	appl, err := application_core_entities.NewApplicationEntity(applData)

	assert.NotNil(t, err)
	assert.Nil(t, appl)
	assert.Equal(t, err.Error(), "please provide application log movement record")
}
