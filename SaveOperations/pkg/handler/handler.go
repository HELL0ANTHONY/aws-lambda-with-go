package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/models"
)

type Processor interface {
	Process(events.APIGatewayProxyRequest) error
}

type Handler struct {
	p Processor
}

func New(p Processor) Handler {
	return Handler{
		p,
	}
}

func (h Handler) Handle(
	ctx context.Context,
	e events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	var id string = lc.AwsRequestID
	if err := h.p.Process(e); err != nil {
		// log.Printf("<start> <Handler> Received request with [%v]", req.QueryStringParameters)
		return models.ResponseError(
			fmt.Sprintf("it was not possible to store the data sent: %s", err.Error()),
			id,
		)
	}

	const message = "Successful Operation"
	return events.APIGatewayProxyResponse{
		Headers:    models.CORSHeaders("*"),
		StatusCode: http.StatusCreated,
		Body: string(
			fmt.Sprintf(`{"message": %q}`, message),
		),
	}, nil
}
