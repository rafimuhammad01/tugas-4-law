package dto

type ReadRequest struct {
	NPM           string `json:"npm"`
	TransactionID int    `json:"transaction_id"`
}

type TransactionReadResponse struct {
	TransactionID int          `json:"transaction_id"`
	Student       ReadResponse `json:"student"`
}

type ReadResponse struct {
	ID   int    `json:"id"`
	NPM  string `json:"npm"`
	Name string `json:"name"`
}
