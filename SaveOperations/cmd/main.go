package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/internal/processor"
	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/internal/repository"
	"github.com/HELL0ANTHONY/aws-lambdas-with-golang/SaveOperations/pkg/handler"
)

func main() {
	r := repository.New()
	p := processor.New(r)
	h := handler.New(p)

	lambda.Start(h.Handle)
}
