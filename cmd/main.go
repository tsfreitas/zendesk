package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/melvi-co/zendesk-integation/internal/handlers"
)

var handler = handlers.GenerateSQSHandler()

func main() {
	// environment configurations and other AWS stuffs
	lambda.Start(handler.HandlerEvent)
}
