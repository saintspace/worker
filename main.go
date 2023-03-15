package main

import (
	"context"
	"sync"

	"worker/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var eventRouter *app.EventRouter

func init() {
	// TODO: construct eventRouter with dependency injection
}

func main() {
	lambda.Start(handler)
}

type LambdaResponse struct {
	BatchItemFailures []events.SQSBatchItemFailure `json:"batchItemFailures"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) (LambdaResponse, error) {
	var response LambdaResponse
	response.BatchItemFailures = []events.SQSBatchItemFailure{}

	var wg sync.WaitGroup
	failureChan := make(chan events.SQSBatchItemFailure, len(sqsEvent.Records))

	for _, sqsMessage := range sqsEvent.Records {
		wg.Add(1)
		go func(msg events.SQSMessage) {
			defer wg.Done()
			err := eventRouter.ProcessEvent(msg.Body)
			if err != nil {
				failureChan <- events.SQSBatchItemFailure{
					ItemIdentifier: msg.MessageId,
				}
			}
		}(sqsMessage)
	}

	wg.Wait()
	close(failureChan)

	for failure := range failureChan {
		response.BatchItemFailures = append(response.BatchItemFailures, failure)
	}

	return response, nil
}
