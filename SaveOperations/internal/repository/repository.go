package repository

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/models"
	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/utils"
)

type Transaction interface {
	Save(*[]models.Operation, *string) error
	WriteItems(*[]models.Record, *sync.WaitGroup, chan error)
}

type Repository struct {
	svc       *dynamodb.DynamoDB
	semaphore chan struct{}
}

const concurrencyLimit = 5

func New() Repository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	return Repository{
		svc:       svc,
		semaphore: make(chan struct{}, concurrencyLimit),
	}
}

func (r Repository) WriteItems(
	operations *[]models.Record,
	wg *sync.WaitGroup,
	errCh chan error,
) {
	const env = "TABLE_NAME"
	tableName, exists := os.LookupEnv(env)
	if !exists || tableName == "" {
		errCh <- fmt.Errorf("%q is not configured correctly", env)
		return
	}
	transacItems := []*dynamodb.TransactWriteItem{}
	for _, op := range *operations {
		operation, err := dynamodbattribute.MarshalMap(op)
		if err != nil {
			errCh <- err
			return
		}
		transacItems = append(transacItems, &dynamodb.TransactWriteItem{
			Put: &dynamodb.Put{
				TableName: aws.String(tableName),
				Item:      operation,
			},
		})
	}
	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transacItems}
	if err := transaction.Validate(); err != nil {
		errCh <- err
		return
	}
	_, err := r.svc.TransactWriteItems(transaction)
	errCh <- err
}

func (r Repository) Save(operations *[]models.Operation, email *string) error {
	record, err := utils.AddMetadata(operations, email)
	if err != nil {
		return fmt.Errorf("an error has occurred while trying to add the metadata: %w", err)
	}

	const chunkSize = 20
	opsChunk := utils.Chunk(record, chunkSize)

	var wg sync.WaitGroup
	errCh := make(chan error, len(opsChunk))
	for _, ops := range opsChunk {
		r.semaphore <- struct{}{}
		wg.Add(1)
		go func(ops []models.Record) {
			defer func() {
				wg.Done()
			}()
			r.WriteItems(&ops, &wg, errCh)
		}(ops)
	}
	wg.Wait()
	close(errCh)

	for e := range errCh {
		if e != nil && err == nil {
			err = e
		}
	}
	return err
}
