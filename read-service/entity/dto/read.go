package dto

type ReadRequest struct {
	NPM string `json:"npm"`
}

type ReadResponse struct {
	ID   int    `json:"id"`
	NPM  string `json:"npm"`
	Name string `json:"name"`
}
