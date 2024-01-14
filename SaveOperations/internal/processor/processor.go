package processor

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"golang.org/x/exp/slog"

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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	email, err := utils.Email(e.Headers["Authorization"])
	if err != nil {
		slog.Error(
			"an error occurred while trying to decode the token",
			slog.String("error=%q", err.Error()),
		)
		return err
	}

	request := models.Request{}
	if err := json.Unmarshal([]byte(e.Body), &request); err != nil {
		slog.Error(
			"an error has occurred while trying to unmarshal the request",
			slog.String("error=%q", err.Error()),
		)
		return err
	}

	return p.r.Save(&request.Data, email)
}
