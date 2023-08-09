package utils

import (
	"errors"
	"fmt"
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

	timestamp, err := Time(func(t time.Time) any {
		return t.Format("2006-01-02T15:04:05")
	})
	if err != nil {
		return []models.Record{}, fmt.Errorf("failed to get timestamp value: %w", err)
	}

	r := []models.Record{}
	for _, op := range *operations {
		uuid := uuid.NewV4().String()
		nano, err := Time(func(t time.Time) any {
			return t.Nanosecond()
		})
		if err != nil {
			return []models.Record{}, fmt.Errorf("failed to get nanosecond value: %w", err)
		}

		op.InternalNumber = "WS" + strconv.Itoa(nano.(int))
		record := models.Record{
			UUID:            uuid,
			OperationStatus: "pendiente",
			Operation:       op,
			EmailCreatedBy:  *email,
			EmailUpdatedBy:  *email,
			Attempts:        0,
			CreatedAt:       timestamp.(string),
			UpdatedAt:       timestamp.(string),
		}
		r = append(r, record)
	}

	return r, nil
}
