package processor

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/utils"
)

type Processor struct{}

func New() Processor {
	return Processor{}
}

func (p Processor) Process(e events.APIGatewayProxyRequest) error {
	email, err := utils.Email(e.Headers["Authorization"])
	if err != nil {
		return fmt.Errorf("an error occurred while trying to decode the token: %w", err)
	}
	log.Println("email", *email)

	return nil
}
