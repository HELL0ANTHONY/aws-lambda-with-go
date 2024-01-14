package models

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

const (
	allowHeaders = "Access-Control-Allow-Headers"
	allowMethods = "Access-Control-Allow-Methods"
	allowOrigin  = "Access-Control-Allow-Origin"
	contentType  = "Content-Type"
)

func CORSHeaders() map[string]string {
	domain := os.Getenv("DOMAIN")
	return map[string]string{
		allowHeaders: "Access-Control-Allow-Origin, Access-Control-Allow-Methods, Content-Type",
		allowMethods: "OPTIONS, POST",
		allowOrigin:  domain,
		contentType:  "application/json",
	}
}

func ResponseError(m, id string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    CORSHeaders(),
		StatusCode: http.StatusInternalServerError,
		Body:       string(fmt.Sprintf(`{"message": %q, "operation_id": %q}`, m, id)),
	}, nil
}
