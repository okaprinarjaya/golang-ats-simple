package application_core_entities_tests_fixtures

import (
	"time"

	constants "gitlab.com/okaprinarjaya.wartek/ats-simple/modules"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_valueobjects "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

func CreateApplicationData() application_core_dto.ApplicationBasicDTO {
	applicationId := "application-20220501-001"
	jobId := "job-001"
	applicantId := "applicant-001"
	applicationCreatedAt := time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC)
	applicationCreatedBy := "Recruiter"

	return application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        applicationId,
			CreatedAt: applicationCreatedAt,
			CreatedBy: applicationCreatedBy,
		},
		ApplicantId:               applicantId,
		JobId:                     jobId,
		CurrentHiringStepSequence: 1,
		Applicant: application_core_valueobjects.ApplicantVO{
			ApplicantCompleteName:   "Budi",
			ApplicantGender:         "MALE",
			ApplicantDateOfBirth:    time.Date(1987, time.March, 20, 0, 0, 0, 0, time.Local),
			ApplicantAddress:        "Jln. Nama jalan tempat tinggal si applicant, RT 001 RW 002, Kelurahan, Kecamatan, Kabupaten, Provinsi",
			ApplicantProfilePhoto:   "https://img.com/?pic=12qwaszxxzsawq21",
			ApplicantProfileSummary: "Detail-oriented Civil Engineer and PE with 5 years of experience and a zest for solving complex problems. Seeking to use proven project management and design skills to improve quality, cost and time metrics for NEP Engineering.",
			ApplicantNationality:    "WNI",
			ApplicantCountryId:      "001",
			ApplicantCountryName:    "Indonesia",
			ApplicantCityId:         "001-001-002",
			ApplicantCityName:       "Jakarta Selatan",
		},
		Job: application_core_valueobjects.JobVO{
			JobName:           "Civil Engineering",
			JobAdmStatus:      "FULL_TIME",
			JobDepartmentId:   "DepartmentId123",
			JobDepartmentName: "Department Name ABC",
			JobCountryId:      "001",
			JobCountryName:    "Indonesia",
			JobCityId:         "001-001-001",
			JobCityName:       "Jakarta Pusat",
		},
	}
}

func CreateApplicationData_InitialCvSubmissionInProgress() application_core_dto.ApplicationBasicDTO {
	applicationId := "application-20220501-001"
	jobId := "job-001"
	applicationLogCreatedBy := "Recruiter"

	applData := CreateApplicationData()
	applData.CurrentHiringStepSequence = 1
	applData.ApplicationLogs = append(applData.ApplicationLogs, application_core_dto.ApplicationLogBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        "appl-log-001",
			CreatedAt: time.Date(2022, time.May, 1, 10, 0, 0, 0, time.UTC),
			CreatedBy: applicationLogCreatedBy,
		},
		ApplicationId:      applicationId,
		JobId:              jobId,
		HiringStepType:     constants.HIRING_STEP_TYPE_CV_SUBMISSION,
		HiringStepSequence: 1,
		HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
	})

	return applData
}

func CreateApplicationData_WithTwoStepMovementLog() application_core_dto.ApplicationBasicDTO {
	applicationId := "application-20220501-001"
	jobId := "job-001"
	applicationLogCreatedBy := "Recruiter"

	applData := CreateApplicationData()
	applData.CurrentHiringStepSequence = 1
	applData.ApplicationLogs = []application_core_dto.ApplicationLogBasicDTO{
		{
			BaseRecord: core_shared.BaseDTO{
				Id:        "appl-log-001",
				CreatedAt: time.Date(2022, time.May, 1, 10, 0, 0, 0, time.UTC),
				CreatedBy: applicationLogCreatedBy,
			},
			ApplicationId:      applicationId,
			JobId:              jobId,
			HiringStepType:     constants.HIRING_STEP_TYPE_CV_SUBMISSION,
			HiringStepSequence: 1,
			HiringStepStatus:   constants.APPL_STEP_STATUS_PASSED,
		},
		{
			BaseRecord: core_shared.BaseDTO{
				Id:        "appl-log-002",
				CreatedAt: time.Date(2022, time.May, 2, 10, 0, 0, 0, time.UTC),
				CreatedBy: applicationLogCreatedBy,
			},
			ApplicationId:      applicationId,
			JobId:              jobId,
			HiringStepType:     constants.HIRING_STEP_TYPE_INTERVIEW,
			HiringStepSequence: 2,
			HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
		},
	}

	return applData
}
