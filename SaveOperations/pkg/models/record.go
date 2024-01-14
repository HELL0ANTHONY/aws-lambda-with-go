package models

type Record struct {
	CreatedAt       string    `json:"createdAt"       dynamodbav:"createdAt"`
	EmailCreatedBy  string    `json:"emailCreatedBy"  dynamodbav:"emailCreatedBy"`
	EmailUpdatedBy  string    `json:"emailUpdatedBy"  dynamodbav:"emailUpdatedBy"`
	OperationStatus string    `json:"operationStatus" dynamodbav:"operationStatus"`
	UUID            string    `json:"uuid"            dynamodbav:"uuid"`
	UpdatedAt       string    `json:"updatedAt"       dynamodbav:"updatedAt"`
	Operation       Operation `json:"operation"       dynamodbav:"operation"`
	Attempts        int       `json:"attempts"        dynamodbav:"attempts"`
}
