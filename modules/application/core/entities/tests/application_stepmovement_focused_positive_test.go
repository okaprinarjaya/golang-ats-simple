package application_core_entities_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_core_entities_tests_fixtures "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities/tests/fixtures"
)

func TestMoveApplicationToNextStep_Positive(t *testing.T) {
	appl, err := application_core_entities.NewApplicationEntity(application_core_entities_tests_fixtures.GetOneApplicationData())

	// Make sure successfully create new application data
	assert.Nil(t, err)

	// Make sure initial created new application data must in step 1
	assert.Equal(t, 1, appl.CurrentStepSequence())

	err = appl.MoveToNextStep(2, application_core_entities.APPL_STEP_STATUS_IN_PROGRESS)

	assert.Nil(t, err)
	assert.Equal(t, 2, appl.CurrentStepSequence())
	assert.Equal(t, 1, len(appl.ApplicationLogs()))
	assert.Equal(t, "IN_PROGRESS", appl.ApplicationLogs()[0].StepStatus())
}
