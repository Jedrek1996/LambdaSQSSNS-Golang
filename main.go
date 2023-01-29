package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

var ctx context.Context

func init() {
	//Inject clients into context and start lambda with those context

	ctx = context.Background() // Empty context

	//Inject aws Client
}

// Takes in the event and process
func Handler(ctx context.Context, event interface{}) {
	// sqsClient := GetSQSClientFromCtx(ctx)
}

func main() {
	lambda.StartWithContext(Handler)
}
