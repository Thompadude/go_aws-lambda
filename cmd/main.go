package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda" // <-- the Lambda programming model for Go
)

type MyEvent struct {
	Name string `json:"name"`
}

// the Lambda handler signature, includes the code which will be executed
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	if name.Name == "asdf" {
		return fmt.Sprint("٩(̾●̮̮̃̾•̃̾)۶"), nil
	}

	if len(name.Name) < 0 {
		return fmt.Sprint("Hello nameless person!"), nil
	}

	return fmt.Sprintf("Hello %s!", name.Name), nil
}

// The entry point that executes your Lambda function code. This is required.
func main() {
	lambda.Start(HandleRequest)
}
