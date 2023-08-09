package repository

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/models"
)

type Repository struct {
	svc *dynamodb.DynamoDB
}

func New() Repository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	return Repository{svc}
}

// Probar enviando *[]models.operations{} (vacío)
func (r Repository) Save(operations *[]models.Operation, email *string) error {
	log.Println("email", *email)
	log.Println("operations", operations)
	return nil
}
