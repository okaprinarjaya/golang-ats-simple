package application_core_entities_tests_fixtures

import (
	"time"

	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_valueobjects "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

func CreateApplicationData() application_core_dto.ApplicationBasicDTO {
	id := "application-20220501-001"
	createdAt := time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC)
	createdBy := "Recruiter"

	return application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        id,
			CreatedAt: createdAt,
			CreatedBy: createdBy,
		},
		ApplicantId: "applicant-001",
		JobId:       "job-001",
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
			JobSAdmtatus:      "FULL_TIME",
			JobDepartmentId:   "DepartmentId123",
			JobDepartmentName: "Department Name ABC",
			JobCountryId:      "001",
			JobCountryName:    "Indonesia",
			JobCityId:         "001-001-001",
			JobCityName:       "Jakarta Pusat",
		},
	}
}

func GetOneApplicationData() application_core_dto.ApplicationBasicDTO {
	id := "application-20220501-001"
	createdAt := time.Date(2022, time.May, 1, 6, 0, 0, 0, time.UTC)
	createdBy := "Recruiter"

	return application_core_dto.ApplicationBasicDTO{
		BaseRecord: core_shared.BaseDTO{
			Id:        id,
			CreatedAt: createdAt,
			CreatedBy: createdBy,
		},
		ApplicantId: "applicant-001",
		JobId:       "job-001",
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
			JobSAdmtatus:      "FULL_TIME",
			JobDepartmentId:   "DepartmentId123",
			JobDepartmentName: "Department Name ABC",
			JobCountryId:      "001",
			JobCountryName:    "Indonesia",
			JobCityId:         "001-001-001",
			JobCityName:       "Jakarta Pusat",
		},
	}
}
