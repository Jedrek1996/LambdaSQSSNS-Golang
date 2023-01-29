package main

import (
	"LambdaSNSSQS/Go/pkg/aws"
	"LambdaSNSSQS/Go/pkg/logger"
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var ctx context.Context

func init() {
	log, _ := zap.NewProduction()
	c := aws.New()

	//Inject clients into context and start lambda with those context
	ctx = context.Background() // Empty context

	//Inject logger and aws Client into context
	ctx = logger.Inject(ctx, log)
	ctx = aws.Inject(ctx, c)
}

// Takes in the event and process
func Handler(ctx context.Context, event interface{}) {
	// sqsClient := GetSQSClientFromCtx(ctx)
}

func main() {
	lambda.StartWithContext(ctx, Handler)
}
