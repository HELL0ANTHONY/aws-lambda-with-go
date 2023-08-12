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

	r := []models.Record{}
	for _, op := range *operations {
		uuid := uuid.NewV4().String()
		nano := Time(func(t time.Time) int {
			return t.Nanosecond()
		})

		op.InternalNumber = "WS" + strconv.Itoa(nano)
		record := models.Record{
			UUID:            uuid,
			OperationStatus: "pendiente",
			Operation:       op,
			EmailCreatedBy:  *email,
			EmailUpdatedBy:  *email,
			Attempts:        0,
			CreatedAt:       timestamp,
			UpdatedAt:       timestamp,
		}
		r = append(r, record)
	}

	return r, nil
}
