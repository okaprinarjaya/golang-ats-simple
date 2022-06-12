package application_core_entities_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_core_entities_tests_fixtures "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities/tests/fixtures"
)

func TestCreateApplication_PositiveCase(t *testing.T) {
	applData := application_core_entities_tests_fixtures.CreateApplicationData()
	appl, err := application_core_entities.NewApplicationEntity(applData)

	assert.Nil(t, err)
	assert.Equal(t, 1, appl.CurrentStepSequence())
	assert.Equal(t, 1, len(appl.ApplicationLogs()))

	appLog := &appl.ApplicationLogs()[0]
	assert.Equal(t, appl.Id(), appLog.ApplicationId())
	assert.Equal(t, application_core_entities.APPL_STEP_STATUS_IN_PROGRESS, appLog.StepStatus())
}
