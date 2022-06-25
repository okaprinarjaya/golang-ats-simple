package application_repositories_tests_fixtures

import (
	"time"

	constants "gitlab.com/okaprinarjaya.wartek/ats-simple/modules"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationDataSample_DTO struct {
	ApplicationID               string
	ApplicationCreatedAt        time.Time
	ApplicationCreatedBy        string
	ApplicationCreatedByName    string
	ApplicantID                 string
	JobID                       string
	ApplicationLogID            string
	ApplicationLogCreatedAt     time.Time
	ApplicationLogCreatedBy     string
	ApplicationLogCreatedByName string
}

func ApplicationDataSample1_DTO(dataIdentifier ApplicationDataSample_DTO) application_core_dto.ApplicationBasicDTO {
	applDTO := application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:            dataIdentifier.ApplicationID,
			CreatedAt:     dataIdentifier.ApplicationCreatedAt,
			CreatedBy:     dataIdentifier.ApplicationCreatedBy,
			CreatedByName: dataIdentifier.ApplicationCreatedByName,
		},
		ApplicationLogs: []application_core_dto.ApplicationLogBasicDTO{
			{
				BaseRecord: core_shared.BaseDTO{
					Id:            dataIdentifier.ApplicationLogID,
					CreatedAt:     dataIdentifier.ApplicationLogCreatedAt,
					CreatedBy:     dataIdentifier.ApplicationLogCreatedBy,
					CreatedByName: dataIdentifier.ApplicationLogCreatedByName,
				},
				ApplicationId:      dataIdentifier.ApplicationID,
				JobId:              dataIdentifier.JobID,
				HiringStepType:     constants.HIRING_STEP_TYPE_CV_SUBMISSION,
				HiringStepSequence: 1,
				HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
				UserType:           "APPLICANT",
			},
		},
		ApplicantId:               dataIdentifier.ApplicantID,
		JobId:                     dataIdentifier.JobID,
		CurrentHiringStepSequence: 1,
		Applicant: application_core_vo.ApplicantVO{
			ApplicantCompleteName:           "Tono Reverseriansyah",
			ApplicantGender:                 "M",
			ApplicantDateOfBirth:            time.Date(1987, time.March, 20, 0, 0, 0, 0, time.Local),
			ApplicantAddress:                "Jln. Applicant Address No. 001, Applicant City, Applicant Province, Indonesia",
			ApplicantProfilePhoto:           "https://photo.me/tono.png",
			ApplicantProfileSummary:         "I Love my self and i love to feed fish at pool fish",
			ApplicantNationality:            "WNI",
			ApplicantCountryId:              "1",
			ApplicantCountryName:            "Indonesia",
			ApplicantCityId:                 "11",
			ApplicantCityName:               "Yogyakarta",
			ApplicantIsLookingOppty:         true,
			ApplicantEducationLast:          "Education Last",
			ApplicantTotalYearsXp:           3,
			ApplicantJobLevelLast:           "Level Last",
			ApplicantWillingWorkOverseas:    false,
			ApplicantExpectedSalary:         10000000,
			ApplicantExpectedSalaryCurrency: "IDR",
		},
		Job: application_core_vo.JobVO{
			JobName:           "Engineering Manager",
			JobAdmStatus:      "FULL_TIME",
			JobDepartmentId:   "1",
			JobDepartmentName: "Tech Engineering",
			JobCountryId:      "1",
			JobCountryName:    "Indonesia",
			JobCityId:         "11",
			JobCityName:       "Yogyakarta",
		},
	}

	return applDTO
}
