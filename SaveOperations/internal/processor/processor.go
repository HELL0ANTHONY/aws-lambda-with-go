package processor

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/models"
	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/utils"
)

type Repository interface {
	Save(*[]models.Operation, *string) error
}

type Processor struct {
	r Repository
}

func New(r Repository) Processor {
	return Processor{
		r,
	}
}

func (p Processor) Process(e events.APIGatewayProxyRequest) error {
	email, err := utils.Email(e.Headers["Authorization"])
	if err != nil {
		return fmt.Errorf("an error occurred while trying to decode the token: %w", err)
	}

	request := models.Request{}
	if err := json.Unmarshal([]byte(e.Body), &request); err != nil {
		return fmt.Errorf("an error has occurred while trying to unmarshal the request: %w", err)
	}

	return p.r.Save(&request.Data, email)
}
