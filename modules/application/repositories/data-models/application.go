package application_repositories_datamodels

import (
	"database/sql"
	"time"
)

type Application struct {
	ID                        string `gorm:"primaryKey"`
	ApplicantId               string `gorm:"column:applicant_id"`
	JobId                     string `gorm:"column:job_id"`
	CurrentHiringStepSequence int    `gorm:"column:current_hiring_step_sequence"`
	IsRejected                bool   `gorm:"column:is_rejected"`
	IsCancelled               bool   `gorm:"column:is_cancelled"`
	IsWithdrawed              bool   `gorm:"column:is_withdrawed"`
	IsOffered                 bool   `gorm:"column:is_offered"`
	IsHired                   bool   `gorm:"column:is_hired"`

	ApplicantCompleteName           string    `gorm:"column:applicant_complete_name"`
	ApplicantGender                 string    `gorm:"column:applicant_gender"`
	ApplicantDateOfBirth            time.Time `gorm:"column:applicant_date_of_birth"`
	ApplicantAddress                string    `gorm:"column:applicant_address"`
	ApplicantProfilePhoto           string    `gorm:"column:applicant_profile_photo"`
	ApplicantProfileSummary         string    `gorm:"column:applicant_profile_summary"`
	ApplicantNationality            string    `gorm:"column:applicant_nationality"`
	ApplicantCountryId              string    `gorm:"column:applicant_country_id"`
	ApplicantCountryName            string    `gorm:"column:applicant_country_name"`
	ApplicantCityId                 string    `gorm:"column:applicant_city_id"`
	ApplicantCityName               string    `gorm:"column:applicant_city_name"`
	ApplicantIsLookingOppty         bool      `gorm:"column:applicant_is_looking_oppty"`
	ApplicantEducationLast          string    `gorm:"column:applicant_education_last"`
	ApplicantTotalYearsXp           int       `gorm:"column:applicant_total_years_xp"`
	ApplicantJobLevelLast           string    `gorm:"column:applicant_job_level_last"`
	ApplicantWillingWorkOverseas    bool      `gorm:"column:applicant_willing_work_overseas"`
	ApplicantExpectedSalary         int       `gorm:"column:applicant_expected_salary"`
	ApplicantExpectedSalaryCurrency string    `gorm:"column:applicant_expected_salary_currency"`

	JobName           string `gorm:"column:job_name"`
	JobAdmStatus      string `gorm:"column:job_adm_status"`
	JobDepartmentId   string `gorm:"column:job_department_id"`
	JobDepartmentName string `gorm:"column:job_department_name"`
	JobCountryId      string `gorm:"column:job_country_id"`
	JobCountryName    string `gorm:"column:job_country_name"`
	JobCityId         string `gorm:"column:job_city_id"`
	JobCityName       string `gorm:"column:job_city_name"`

	CreatedAt       time.Time        `gorm:"column:created_at"`
	UpdatedAt       sql.NullTime     `gorm:"column:updated_at"`
	DeletedAt       sql.NullTime     `gorm:"column:deleted_at"`
	CreatedBy       string           `gorm:"column:created_by"`
	CreatedByName   string           `gorm:"column:created_by_name"`
	UpdatedBy       sql.NullString   `gorm:"column:updated_by"`
	UpdatedByName   sql.NullString   `gorm:"column:updated_by_name"`
	DeletedBy       sql.NullString   `gorm:"column:deleted_by"`
	DeletedByName   sql.NullString   `gorm:"column:deleted_by_name"`
	ApplicationLogs []ApplicationLog `gorm:"foreignKey:ApplicationId"`
}
