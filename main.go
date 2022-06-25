package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	infrastructure_database "gitlab.com/okaprinarjaya.wartek/ats-simple/infrastructure/database"
	constants "gitlab.com/okaprinarjaya.wartek/ats-simple/modules"
	application_core_dto "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/dto"
	application_core_entities "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/entities"
	application_core_vo "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/core/value-objects"
	application_repositories "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/application/repositories"
	core_shared "gitlab.com/okaprinarjaya.wartek/ats-simple/modules/core-shared"
)

func main() {
	dbconn := infrastructure_database.PgConnection()
	var repo application_repositories.IApplicationRepository = application_repositories.NewApplicationRepositoryPostgreSql(dbconn)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		ksuidAppl := "application-id-123"
		ksuidApplLog := "application-log-id-123"
		jobId := "job-id-123"
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

					ApplicationId:      ksuidAppl,
					JobId:              jobId,
					HiringStepType:     constants.HIRING_STEP_TYPE_CV_SUBMISSION,
					HiringStepSequence: 1,
					HiringStepStatus:   constants.APPL_STEP_STATUS_IN_PROGRESS,
					UserType:           "APPLICANT",
				},
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
				JobAdmStatus:      "FULL_TIME",
				JobDepartmentId:   "DeptId123",
				JobDepartmentName: "Department Name ABC",
				JobCountryId:      "CountryId123",
				JobCountryName:    "Country Name ABC",
				JobCityId:         "CityId123",
				JobCityName:       "City Name ABC",
			},
		}

		appl, _ := application_core_entities.NewApplicationEntity(applDTO)

		applLogBaseRecordDTO := core_shared.BaseDTO{
			Id: "123abc",
		}

		appl.MoveToNextStep(applLogBaseRecordDTO, 2, "INTERVIEW", "RECRUITER")

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
			ApplicationLogs: func() []application_core_dto.ApplicationLogBasicDTO {
				var applLogList []application_core_dto.ApplicationLogBasicDTO
				for _, appLog := range appl.ApplicationLogs() {
					applLogList = append(applLogList, application_core_dto.ApplicationLogBasicDTO{
						BaseRecord: core_shared.BaseDTO{
							Id: appLog.Id(),
						},
						ApplicationId:             appLog.ApplicationId(),
						JobId:                     appLog.JobId(),
						HiringStepType:            appLog.HiringStepType(),
						HiringStepTypeCompletedAt: appLog.HiringStepTypeCompletedAt(),
						HiringStepSequence:        appLog.HiringStepSequence(),
						HiringStepStatus:          appLog.HiringStepStatus(),
					})
				}
				return applLogList
			}(),
			ApplicantId:               appl.ApplicantId(),
			JobId:                     appl.JobId(),
			CurrentHiringStepSequence: appl.CurrentHiringStepSequence(),
			Applicant:                 appl.Applicant(),
			Job:                       appl.Job(),
		})
	})

	r.Post("/applications", func(w http.ResponseWriter, req *http.Request) {
		applicantId := "applicant-id-14323-abc"
		ksuidAppl := "application-id-14323_a"
		ksuidApplLog := "application-log-id-14323__11111111111111111111111111111111111111111111111111111111111111111111111111111"
		jobId := "job-id-12344"
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

		appl, _ := application_core_entities.NewApplicationEntity(applDTO)
		appl.PersistenceStatus = core_shared.NEW
		for i := 0; i < len(appl.ApplicationLogs()); i++ {
			appl.ApplicationLogs()[i].PersistenceStatus = core_shared.NEW
		}

		err := repo.Save(*appl)

		if err != nil {
			fmt.Println(err.Error())
			resp := struct {
				Message string
			}{
				Message: "Responding Error: " + err.Error(),
			}

			render.Status(req, http.StatusInternalServerError)
			render.JSON(w, req, resp)
		} else {
			resp := struct {
				Message string
			}{
				Message: "Responding",
			}

			render.JSON(w, req, resp)
		}

	})

	http.ListenAndServe(":3000", r)
}
