package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabriel-dantas98/golang-lambda-scaffolder/pkg/shouldideploy"
)

type Event struct {
	Message       string `json:"message"`
	ShouldIdeploy bool   `json:"shouldideploy"`
	Timezone      bool   `json:"timezone"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(ctx context.Context, event Event) (Response, error) {
	fmt.Println("incoming event", event)

	message, err := shouldideploy.Request()

	return Response{Message: message}, err
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
