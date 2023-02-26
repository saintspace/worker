package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, sqsMessage := range sqsEvent.Records {
		fmt.Printf("SQS message body: %s\n", sqsMessage.Body)
	}
	return nil
}