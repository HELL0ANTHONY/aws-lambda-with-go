package utils

import (
	"errors"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/models"
)

func AddMetadata(operations *[]models.Operation, email *string) ([]models.Record, error) {
	if email == nil {
		return []models.Record{}, errors.New("email cannot be nil")
	}

	if operations == nil {
		return []models.Record{}, errors.New("operations cannot be nil")
	}

	timestamp := Time(func(t time.Time) string {
		return t.Format("2006-01-02T15:04:05")
	})

	numberOfOperations := len(*operations)
	records := make([]models.Record, 0, numberOfOperations)
	for i := 0; i < numberOfOperations; i++ {
		uuid := uuid.NewV4().String()
		nano := Time(func(t time.Time) int {
			return t.Nanosecond()
		})

		op := (*operations)[i]
		op.InternalNumber = "WS" + strconv.Itoa(nano)
		record := models.Record{
			Attempts:        0,
			CreatedAt:       timestamp,
			EmailCreatedBy:  *email,
			EmailUpdatedBy:  *email,
			Operation:       op,
			OperationStatus: "pendiente",
			UUID:            uuid,
			UpdatedAt:       timestamp,
		}
		records = append(records, record)
	}

	return records, nil
}
