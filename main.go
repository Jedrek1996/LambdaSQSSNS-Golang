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

	ctx = context.Background()    // Empty context
	log, _ := zap.NewProduction() // To catch errors etc
	c, err := aws.New()

	if err != nil {
		log.Fatal("Cannot set up AWS Clients")
	}

	//Inject clients into context and start lambda with those context
	//Inject Logger and aws Client into context
	ctx = logger.Inject(ctx, log)
	ctx = aws.Inject(ctx, c)
}

// Takes in the event and process
func Handler(ctx context.Context, event interface{}) {

	log := logger.GetLoggerFromContext(ctx)
	awsConnection := aws.GetConnectionFromContext(ctx)

	err := awsConnection.SendSQSMessage(ctx, "SQS HELLO")
	if err != nil {
		log.Error("Can't send SQS!", zap.Error(err))
	}

	err = awsConnection.PublishSNSMessage(ctx, "SNS HEY!")
	if err != nil {
		log.Error("Can't send SNS", zap.Error(err))
	}

}

func main() {
	lambda.StartWithContext(ctx, Handler)
}
