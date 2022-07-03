package application_core_entities_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_core_entities_tests_fixtures "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities/tests/fixtures"
)

func TestCreateApplication_InitialCvSubmissionInProgress_NegativeCase(t *testing.T) {
	applSampleData := application_core_entities_tests_fixtures.ApplicationDataSample_DTO{
		ApplicationID:            "application-id-001",
		ApplicantID:              "applicant-id-001",
		JobID:                    "job-id-001",
		ApplicationCreatedAt:     time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC),
		ApplicationCreatedBy:     "created-by-001",
		ApplicationCreatedByName: "Thomas Shelby the Applicant",
	}
	applData := application_core_entities_tests_fixtures.CreateApplicationData(applSampleData)
	appl, err := application_core_entities.NewApplicationEntity(applData)

	assert.NotNil(t, err)
	assert.Nil(t, appl)
	assert.Equal(t, err.Error(), "please provide application log movement record")
}
