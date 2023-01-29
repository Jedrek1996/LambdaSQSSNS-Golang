package aws

import (
	"LambdaSNSSQS/Go/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/zap"
)

type Connection struct {
	sqs *sqs.Client
	sns *sns.Client
}

func New() (*Connection, error) {

	//To load in aws configs/ credentials (set up in CLI)
	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		return nil, err
	}

	return &Connection{
		sqs: sqs.NewFromConfig(cfg), //Set up a new client with the configs
		sns: sns.NewFromConfig(cfg),
	}, nil
}

func (c *Connection) SendSQSMessage(ctx context.Context, message string) error {
	log := logger.GetLoggerFromContext(ctx)

	//Standard &sqs.SendMessageInput struct
	output, err := c.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    new(string),
	})

	log.Info("SQS Message sent:", zap.Any("Output", output))

	return err
}

func (c *Connection) PublishSNSMessage(ctx context.Context, message string) error {
	log := logger.GetLoggerFromContext(ctx)

	//Standard &sns.PublishInput struct
	output, err := c.sns.Publish(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: new(string),
	})
	log.Info("Publish SNS message:", zap.Any("Output", output))

	return err
}
