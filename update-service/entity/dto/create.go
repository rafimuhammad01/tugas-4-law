package dto

type CreateStudentRequest struct {
	NPM  string `json:"npm"`
	Name string `json:"name"`
}

type CreateStudentResponse struct {
	Status string `json:"status"`
}
