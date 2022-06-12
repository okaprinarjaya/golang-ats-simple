package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		appl, _ := application_core_entities.NewApplicationEntity(application_core_dto.ApplicationBasicDTO{
			BaseRecord: core_shared.BaseDTO{
				Id: "Id123",
			},
			ApplicantId: "ApplicantIdIsAlsoUserId123",
			JobId:       "JobId123",
			Applicant: application_core_vo.ApplicantVO{
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
			Job: application_core_vo.JobVO{
				JobName:           "Job Name 123",
				JobSAdmtatus:      "FULL_TIME",
				JobDepartmentId:   "DeptId123",
				JobDepartmentName: "Department Name ABC",
				JobCountryId:      "CountryId123",
				JobCountryName:    "Country Name ABC",
				JobCityId:         "CityId123",
				JobCityName:       "City Name ABC",
			},
		})

		appl.MoveToNextStep(2, "INTERVIEW", "Recruiter")

		render.JSON(w, req, application_core_dto.ApplicationBasicDTO{
			BaseRecord: core_shared.BaseDTO{
				Id:        appl.Id(),
				CreatedAt: appl.CreatedAt(),
				UpdatedAt: appl.UpdatedAt(),
				DeletedAt: appl.DeletedAt(),
				CreatedBy: appl.CreatedBy(),
				UpdatedBy: appl.UpdatedBy(),
				DeletedBy: appl.DeletedBy(),
			},
			ApplicantId:         appl.ApplicantId(),
			JobId:               appl.JobId(),
			CurrentStepSequence: appl.CurrentStepSequence(),
			Applicant:           appl.Applicant(),
			Job:                 appl.Job(),
		})
	})

	http.ListenAndServe(":3000", r)
}
