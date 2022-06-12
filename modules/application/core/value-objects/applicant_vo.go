package application_core_valueobjects

import "time"

type ApplicantVO struct {
	ApplicantCompleteName           string
	ApplicantGender                 string
	ApplicantDateOfBirth            time.Time
	ApplicantAddress                string
	ApplicantProfilePhoto           string
	ApplicantProfileSummary         string
	ApplicantNationality            string
	ApplicantCountryId              string
	ApplicantCountryName            string
	ApplicantCityId                 string
	ApplicantCityName               string
	ApplicantIsLookingOppty         bool
	ApplicantEducationLast          string
	ApplicantTotalYearsXp           int
	ApplicantJobLevelLast           string
	ApplicantWillingWorkOverseas    bool
	ApplicantExpectedSalary         int
	ApplicantExpectedSalaryCurrency string
}
