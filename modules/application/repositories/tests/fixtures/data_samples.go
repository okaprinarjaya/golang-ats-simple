package application_repositories_tests_fixtures

import (
	"time"

	constants "gitlab.com/okaprinarjaya.wartek/ats-simple/modules"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

func DataSample1() application_core_dto.ApplicationBasicDTO {
	applicantId := "applicant-id-4"
	ksuidAppl := "application-id-4"
	ksuidApplLog := "application-log-id-4"
	jobId := "job-id-4"
	createdAt := time.Date(2022, time.May, 1, 10, 0, 0, 0, time.UTC)

	applDTO := application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:            ksuidAppl,
			CreatedAt:     createdAt,
			CreatedBy:     "CreatedById123abc321cba",
			CreatedByName: "Oka The Applicant",
		},
		ApplicationLogs: []application_core_dto.ApplicationLogBasicDTO{
			{
				BaseRecord: core_shared.BaseDTO{
					Id:            ksuidApplLog,
					CreatedAt:     createdAt,
					CreatedBy:     "CreatedById123abc321cba",
					CreatedByName: "Oka The Applicant",
				},
				ApplicationId:                ksuidAppl,
				JobId:                        jobId,
				HiringStepType:               constants.HIRING_STEP_TYPE_CV_SUBMISSION,
				HiringStepStatusClosedBy:     "Closed By ID 123",
				HiringStepStatusClosedByName: "Closed By Name - Tono",
				HiringStepSequence:           1,
				HiringStepStatus:             constants.APPL_STEP_STATUS_IN_PROGRESS,
				UserType:                     "APPLICANT",
			},
		},
		ApplicantId:               applicantId,
		JobId:                     jobId,
		CurrentHiringStepSequence: 1,
		Applicant: application_core_vo.ApplicantVO{
			ApplicantCompleteName:           "Applicant Complete Name",
			ApplicantGender:                 "M",
			ApplicantDateOfBirth:            time.Date(1987, time.March, 20, 0, 0, 0, 0, time.Local),
			ApplicantAddress:                "Jln. Applicant Address No. 001, Applicant City, Applicant Province, Indonesia",
			ApplicantProfilePhoto:           "https://photo.me/oka.png",
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
