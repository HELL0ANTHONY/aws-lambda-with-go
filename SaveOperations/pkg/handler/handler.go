package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"golang.org/x/exp/slog"

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
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	slog.Info("starting", slog.String("data=%q", req.Body))

	lc, _ := lambdacontext.FromContext(ctx)
	id := lc.AwsRequestID

	if err := h.p.Process(req); err != nil {
		slog.Error("it was not possible to store the data sent", slog.String("error=", err.Error()))
		return models.ResponseError(
			err.Error(),
			id,
		)
	}

	const message = "Successful Operation"
	return events.APIGatewayProxyResponse{
		Headers:    models.CORSHeaders(),
		StatusCode: http.StatusCreated,
		Body: string(
			fmt.Sprintf(`{"message": %q}`, message),
		),
	}, nil
}
