package application_core_valueobjects

type JobVO struct {
	JobName           string `json:"jobName"`
	JobSAdmtatus      string `json:"jobAdmStatus"`
	JobDepartmentId   string `json:"jobDepartmentId"`
	JobDepartmentName string `json:"jobDepartmentName"`
	JobCountryId      string `json:"jobCountryId"`
	JobCountryName    string `json:"jobCountryName"`
	JobCityId         string `json:"jobCityId"`
	JobCityName       string `json:"jobCityName"`
}
