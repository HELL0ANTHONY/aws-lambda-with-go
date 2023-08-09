package models

type Record struct {
	Attempts        int       `json:"attempts"`
	CreatedAt       string    `json:"createdAt"`
	EmailCreatedBy  string    `json:"emailCreatedBy"`
	EmailUpdatedBy  string    `json:"emailUpdatedBy"`
	Operation       Operation `json:"operation"`
	OperationStatus string    `json:"operationStatus"`
	UUID            string    `json:"uuid"`
	UpdatedAt       string    `json:"updatedAt"`
}
