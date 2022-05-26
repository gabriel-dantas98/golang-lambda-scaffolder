package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gabriel-dantas98/golang-lambda-scaffolder/handler"
)

func main() {
	lambda.Start(handler.HandleLambdaEvent)
}
