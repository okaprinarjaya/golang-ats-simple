package application_core_entities_tests_fixtures

import (
	"time"

	constants "gitlab.com/okaprinarjaya.wartek/ats-simple/modules"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_valueobjects "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

type ApplicationDataSample_DTO struct {
	CurrentHiringStepSequence int
	ApplicationID             string
	ApplicationCreatedAt      time.Time
	ApplicationCreatedBy      string
	ApplicationCreatedByName  string
	ApplicantID               string
	JobID                     string
}

func CreateApplicationData(applSampleData ApplicationDataSample_DTO) application_core_dto.ApplicationBasicDTO {
	return application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:            applSampleData.ApplicationID,
			CreatedAt:     applSampleData.ApplicationCreatedAt,
			CreatedBy:     applSampleData.ApplicationCreatedBy,
			CreatedByName: applSampleData.ApplicationCreatedByName,
		},
		ApplicantId:               applSampleData.ApplicantID,
		JobId:                     applSampleData.JobID,
		CurrentHiringStepSequence: applSampleData.CurrentHiringStepSequence,
		Applicant: application_core_valueobjects.ApplicantVO{
			ApplicantCompleteName:           applSampleData.ApplicationCreatedByName,
			ApplicantGender:                 "MALE",
			ApplicantDateOfBirth:            time.Date(1987, time.March, 20, 0, 0, 0, 0, time.Local),
			ApplicantAddress:                "Jln. Nama jalan tempat tinggal si applicant, RT 001 RW 002, Kelurahan, Kecamatan, Kabupaten, Provinsi",
			ApplicantProfilePhoto:           "https://img.com/?pic=12qwaszxxzsawq21",
			ApplicantProfileSummary:         "Detail-oriented Civil Engineer and PE with 5 years of experience and a zest for solving complex problems. Seeking to use proven project management and design skills to improve quality, cost and time metrics for NEP Engineering.",
			ApplicantNationality:            "WNI",
			ApplicantCountryId:              "001",
			ApplicantCountryName:            "Indonesia",
			ApplicantCityId:                 "001-001-002",
			ApplicantCityName:               "Jakarta Selatan",
			ApplicantIsLookingOppty:         true,
			ApplicantEducationLast:          "Education last",
			ApplicantTotalYearsXp:           3,
			ApplicantJobLevelLast:           "Job level last",
			ApplicantWillingWorkOverseas:    false,
			ApplicantExpectedSalary:         13000000,
			ApplicantExpectedSalaryCurrency: "IDR",
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

func CreateApplicationData_CVSubmission_InProgress() application_core_dto.ApplicationBasicDTO {
	applSampleData := ApplicationDataSample_DTO{
		CurrentHiringStepSequence: 1,
		ApplicationID:             "application-id-001",
		ApplicantID:               "applicant-id-001",
		JobID:                     "job-id-001",
		ApplicationCreatedAt:      time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC),
		ApplicationCreatedBy:      "applicant-id-001",
		ApplicationCreatedByName:  "Thomas Shelby the Applicant",
	}

	applData := CreateApplicationData(applSampleData)
	applData.ApplicationLogs = append(applData.ApplicationLogs, application_core_dto.ApplicationLogBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:            "application-log-id-001",
			CreatedAt:     time.Date(2022, time.May, 1, 10, 0, 0, 0, time.UTC),
			CreatedBy:     "applicant-id-001",
			CreatedByName: "Thomas Shelby the Applicant",
		},
		ApplicationId:      applSampleData.ApplicationID,
		JobId:              applSampleData.JobID,
		HiringStepType:     constants.HIRING_STEP_TYPE_CV_SUBMISSION,
		HiringStepSequence: 1,
		HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
		UserType:           "APPLICANT",
	})

	return applData
}

func CreateApplicationData_Step2_Interview_InProgress() application_core_dto.ApplicationBasicDTO {
	applSampleData := ApplicationDataSample_DTO{
		CurrentHiringStepSequence: 2,
		ApplicationID:             "application-id-001",
		ApplicantID:               "applicant-id-001",
		JobID:                     "job-id-001",
		ApplicationCreatedAt:      time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC),
		ApplicationCreatedBy:      "applicant-id-001",
		ApplicationCreatedByName:  "Thomas Shelby the Applicant",
	}

	applData := CreateApplicationData(applSampleData)
	applData.ApplicationLogs = []application_core_dto.ApplicationLogBasicDTO{
		{
			BaseRecord: core_shared.BaseDTO{
				Id:            "application-log-id-001",
				CreatedAt:     time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC),
				CreatedBy:     "applicant-id-001",
				CreatedByName: "Thomas Shelby the Applicant",
			},
			ApplicationId:                applSampleData.ApplicationID,
			JobId:                        applSampleData.JobID,
			HiringStepType:               constants.HIRING_STEP_TYPE_CV_SUBMISSION,
			HiringStepSequence:           1,
			HiringStepStatus:             constants.APPL_STEP_STATUS_IN_PROGRESS,
			HiringStepStatusClosedAt:     time.Date(2022, time.May, 2, 6, 0, 0, 0, time.UTC),
			HiringStepStatusClosedBy:     "recruiter-id-001",
			HiringStepStatusClosedByName: "The Recruiter Name",
			UserType:                     "APPLICANT",
		},
		{
			BaseRecord: core_shared.BaseDTO{
				Id:            "application-log-id-002",
				CreatedAt:     time.Date(2022, time.May, 2, 6, 0, 0, 0, time.UTC),
				CreatedBy:     "recruiter-id-001",
				CreatedByName: "The Recruiter Name",
			},
			ApplicationId:      applSampleData.ApplicationID,
			JobId:              applSampleData.JobID,
			HiringStepType:     constants.HIRING_STEP_TYPE_CV_SUBMISSION,
			HiringStepSequence: 1,
			HiringStepStatus:   constants.APPL_STEP_STATUS_PASSED,
			UserType:           "RECRUITER",
		},
		{
			BaseRecord: core_shared.BaseDTO{
				Id:            "application-log-id-003",
				CreatedAt:     time.Date(2022, time.May, 2, 6, 0, 0, 0, time.UTC),
				CreatedBy:     "recruiter-id-001",
				CreatedByName: "The Recruiter Name",
			},
			ApplicationId:      applSampleData.ApplicationID,
			JobId:              applSampleData.JobID,
			HiringStepType:     constants.HIRING_STEP_TYPE_INTERVIEW,
			HiringStepSequence: 2,
			HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
			UserType:           "RECRUITER",
		},
	}

	return applData
}
