package aws

import (
	"context"
)

type AWSKeyType string

var AWSKey AWSKeyType = "AWS"

func Inject(ctx context.Context, c *Connection) context.Context {
	return context.WithValue(ctx, AWSKey, c)
}
