package application_core_valueobjects

import "time"

type ApplicantVO struct {
	ApplicantCompleteName           string    `json:"applicantCompleteName"`
	ApplicantGender                 string    `json:"applicantGender"`
	ApplicantDateOfBirth            time.Time `json:"applicantDateOfBirth"`
	ApplicantAddress                string    `json:"applicantAddress"`
	ApplicantProfilePhoto           string    `json:"applicantProfilePhoto"`
	ApplicantProfileSummary         string    `json:"applicantProfileSummary"`
	ApplicantNationality            string    `json:"applicantNationality"`
	ApplicantCountryId              string    `json:"applicantCountryId"`
	ApplicantCountryName            string    `json:"applicantCountryName"`
	ApplicantCityId                 string    `json:"applicantCityId"`
	ApplicantCityName               string    `json:"applicantCityName"`
	ApplicantIsLookingOppty         bool      `json:"applicantIsLookingOppty"`
	ApplicantEducationLast          string    `json:"applicantEducationLast"`
	ApplicantTotalYearsXp           int       `json:"applicantTotalYearsXp"`
	ApplicantJobLevelLast           string    `json:"applicantJobLevelLast"`
	ApplicantWillingWorkOverseas    bool      `json:"applicantWillingWorkOverseas"`
	ApplicantExpectedSalary         int       `json:"applicantExpectedSalary"`
	ApplicantExpectedSalaryCurrency string    `json:"applicantExpectedSalaryCurrency"`
}
