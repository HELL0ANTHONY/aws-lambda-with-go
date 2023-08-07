package models

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/utils"
)

type Response struct {
	OperationID string `json:"operation_id"`
	Message     string `json:"message"`
}

func ResponseError(m, id string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    utils.CORSHeaders("*"),
		StatusCode: http.StatusInternalServerError,
		Body:       string(fmt.Sprintf(`{"message": %q, "operation_id": %q}`, m, id)),
	}, nil
}
