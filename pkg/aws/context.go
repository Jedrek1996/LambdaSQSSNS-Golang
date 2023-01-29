package aws

import (
	"context"
	"fmt"
)

type AWSKeyType string

var AWSKey AWSKeyType = "AWS"

/*To inject and retrieve the aws connection (from aws.go)*/

func Inject(ctx context.Context, c *Connection) context.Context {
	return context.WithValue(ctx, AWSKey, c)
}

func GetConnectionFromContext(ctx context.Context) *Connection {
	c, ok := ctx.Value(AWSKey).(*Connection)
	//ctx.Value("AWSKey") returns an interface. So .(*myStruct) does a type assertion to convert it to the type *Connection

	if !ok {
		fmt.Println("Connection from Context Error ")
	}

	return c
}
