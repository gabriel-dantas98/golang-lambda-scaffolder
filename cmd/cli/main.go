// cli/main.go
package cli

import (
	"context"
	"fmt"
	"gabriel-dantas98/golang-lambda-scaffolder/cmd/handler"
)

func Exec() {
	fmt.Println("Invoke cli golang")
	fmt.Println(handler.HandleLambdaEvent(context.Background(), handler.Event{}))
}
