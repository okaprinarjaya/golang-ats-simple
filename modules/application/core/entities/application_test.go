package application_core_entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_valueobjects "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

func TestCreateApplication_Success(t *testing.T) {
	id := "appl123"
	createdAt := time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC)
	createdBy := "Created by name ABC"

	applData := application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        id,
			CreatedAt: createdAt,
			CreatedBy: createdBy,
		},
		ApplicantId: "UserId123",
		JobId:       "JobId123",
		Applicant: application_core_valueobjects.ApplicantVO{
			ApplicantCompleteName:   "Applicant Complete Name ABC",
			ApplicantGender:         "",
			ApplicantDateOfBirth:    time.Date(1987, time.March, 20, 0, 0, 0, 0, time.Local),
			ApplicantAddress:        "",
			ApplicantProfilePhoto:   "",
			ApplicantProfileSummary: "",
			ApplicantNationality:    "",
			ApplicantCountryId:      "",
			ApplicantCountryName:    "",
			ApplicantCityId:         "",
			ApplicantCityName:       "",
		},
		Job: application_core_valueobjects.JobVO{
			JobName:           "Job Name 123",
			JobSAdmtatus:      "FULL_TIME",
			JobDepartmentId:   "DepartmentId123",
			JobDepartmentName: "Department Name ABC",
			JobCountryId:      "CountryId123",
			JobCountryName:    "Country Name ABC",
			JobCityId:         "CountryId123",
			JobCityName:       "Country Name ABC",
		},
	}

	appl, err := NewApplicationEntity(applData)

	assert.Nil(t, err)
	assert.Equal(t, 1, appl.CurrentStepSequence())
}

func TestMoveApplicationToNextStep_Success(t *testing.T) {
	id := "appl123"
	createdAt := time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC)
	createdBy := "Created by name ABC"

	applData := application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        id,
			CreatedAt: createdAt,
			CreatedBy: createdBy,
		},
		ApplicantId: "UserId123",
		JobId:       "JobId123",
		Applicant: application_core_valueobjects.ApplicantVO{
			ApplicantGender:         "",
			ApplicantDateOfBirth:    time.Date(1987, time.March, 20, 0, 0, 0, 0, time.Local),
			ApplicantAddress:        "",
			ApplicantProfilePhoto:   "",
			ApplicantProfileSummary: "",
			ApplicantNationality:    "",
			ApplicantCountryId:      "",
			ApplicantCountryName:    "",
			ApplicantCityId:         "",
			ApplicantCityName:       "",
		},
		Job: application_core_valueobjects.JobVO{
			JobName:           "Job Name 123",
			JobSAdmtatus:      "FULL_TIME",
			JobDepartmentId:   "DepartmentId123",
			JobDepartmentName: "Department Name ABC",
			JobCountryId:      "CountryId123",
			JobCountryName:    "Country Name ABC",
			JobCityId:         "CountryId123",
			JobCityName:       "Country Name ABC",
		},
	}

	appl, err := NewApplicationEntity(applData)

	assert.Nil(t, err)
	assert.Equal(t, 1, appl.CurrentStepSequence())

	err = appl.MoveToNextStep(2, "IN_PROGRESS")

	assert.Nil(t, err)
	assert.Equal(t, 2, appl.CurrentStepSequence())
	assert.Equal(t, 1, len(appl.ApplicationLogs()))
	assert.Equal(t, "IN_PROGRESS", appl.ApplicationLogs()[0].StepStatus())
}
