package application_core_entities_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	constants "gitlab.com/okaprinarjaya.wartek/ats-simple/modules"
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_core_entities_tests_fixtures "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities/tests/fixtures"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

func TestStepMovement_MoveFromCVSubmissionToNextStep(t *testing.T) {
	applData := application_core_entities_tests_fixtures.CreateApplicationData_CVSubmission_InProgress()
	appl, err := application_core_entities.NewApplicationEntity(applData)

	assert.Nil(t, err)
	assert.Equal(t, 1, appl.CurrentHiringStepSequence())
	assert.Equal(t, 1, len(appl.ApplicationLogs()))

	appLog := &appl.ApplicationLogs()[0]
	assert.Equal(t, appl.Id(), appLog.ApplicationId())
	assert.Equal(t, constants.APPL_STEP_STATUS_IN_PROGRESS, appLog.HiringStepStatus())

	// Start move
	err = appl.MoveFromCVSubmissionToNextStep(
		2,
		constants.HIRING_STEP_TYPE_INTERVIEW,
		"recruiter-id-001",
		"The Recruiter Name",
		"RECRUITER",
	)

	assert.Nil(t, err)
	assert.Equal(t, 3, len(appl.ApplicationLogs()))
	assert.Equal(t, core_shared.MODIFIED, appl.PersistenceStatus)

	appLog_movement := &appl.ApplicationLogs()[0]
	assert.Equal(t, applData.BaseRecord.Id, appLog_movement.ApplicationId())
	assert.Equal(t, 1, appLog_movement.HiringStepSequence())
	assert.Equal(t, constants.HIRING_STEP_TYPE_CV_SUBMISSION, appLog_movement.HiringStepType())
	assert.Equal(t, constants.APPL_STEP_STATUS_IN_PROGRESS, appLog_movement.HiringStepStatus())
	assert.Equal(t, false, appLog_movement.HiringStepStatusClosedAt().IsZero())
	assert.Equal(t, "recruiter-id-001", appLog_movement.HiringStepStatusClosedBy())
	assert.Equal(t, "The Recruiter Name", appLog_movement.HiringStepStatusClosedByName())
	assert.Equal(t, core_shared.MODIFIED, appLog_movement.PersistenceStatus)

	appLog_movement2 := &appl.ApplicationLogs()[1]
	assert.Equal(t, applData.BaseRecord.Id, appLog_movement2.ApplicationId())
	assert.Equal(t, 1, appLog_movement2.HiringStepSequence())
	assert.Equal(t, constants.HIRING_STEP_TYPE_CV_SUBMISSION, appLog_movement2.HiringStepType())
	assert.Equal(t, constants.APPL_STEP_STATUS_PASSED, appLog_movement2.HiringStepStatus())
	assert.Equal(t, true, appLog_movement2.HiringStepStatusClosedAt().IsZero())
	assert.Equal(t, "", appLog_movement2.HiringStepStatusClosedBy())
	assert.Equal(t, "", appLog_movement2.HiringStepStatusClosedByName())
	assert.Equal(t, core_shared.NEW, appLog_movement2.PersistenceStatus)

	appLog_movement3 := &appl.ApplicationLogs()[2]
	assert.Equal(t, applData.BaseRecord.Id, appLog_movement3.ApplicationId())
	assert.Equal(t, 2, appLog_movement3.HiringStepSequence())
	assert.Equal(t, constants.HIRING_STEP_TYPE_INTERVIEW, appLog_movement3.HiringStepType())
	assert.Equal(t, constants.APPL_STEP_STATUS_IN_PROGRESS, appLog_movement3.HiringStepStatus())
	assert.Equal(t, true, appLog_movement3.HiringStepStatusClosedAt().IsZero())
	assert.Equal(t, "", appLog_movement3.HiringStepStatusClosedBy())
	assert.Equal(t, "", appLog_movement3.HiringStepStatusClosedByName())
	assert.Equal(t, core_shared.NEW, appLog_movement3.PersistenceStatus)
}
