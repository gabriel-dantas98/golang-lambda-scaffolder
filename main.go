package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"gabriel-dantas98/golang-lambda-scaffolder/cmd/handler"
)

func main() {
	lambda.Start(handler.HandleLambdaEvent)
}
