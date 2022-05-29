package handler

import (
	"context"
	"fmt"

	"gabriel-dantas98/golang-lambda-scaffolder/pkg/shouldideploy"
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

	response, err := shouldideploy.Request()

	if err != nil {
		panic(err)
	}

	return Response{Message: response.Message}, nil
}
