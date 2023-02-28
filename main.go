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
		task, err := parseTask(sqsMessage.Body)
		if err != nil {
			return fmt.Errorf("error while parsing task => %v", err.Error())
		}
		err = processTask(task)
		if err != nil {
			return fmt.Errorf("error while processing task => %v", err.Error())
		}
	}
	return nil
}
