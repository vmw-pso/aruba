package data

type Resource struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
	Workgroup      string `json:"workgroup"`
	EmploymentType string `json:"employmentType"`
	Manager        string `json:"manager"`
	Active         bool   `json:"active"`
}
