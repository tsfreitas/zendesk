package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/melvi-co/zendesk-integation/internal/domain/model"
	"github.com/melvi-co/zendesk-integation/internal/domain/selector"
)

type SQSHandlerPort interface {
	HandlerEvent(ctx context.Context, event events.SQSEvent) (string, error)
}
type sqsHandler struct {
	useCaseSelector selector.Selector
}

func newSQSHandler() SQSHandlerPort {
	return &sqsHandler{
		useCaseSelector: selector.NewUseCase(),
	}
}

func (s sqsHandler) HandlerEvent(ctx context.Context, event events.SQSEvent) (string, error) {
	fmt.Printf("Event received. %d messages \n", len(event.Records))

	for _, message := range event.Records {
		payload := model.Payload{}

		errJson := json.Unmarshal([]byte(message.Body), &payload)
		if errJson != nil {
			return "error", errJson
		}

		errUseCase := s.useCaseSelector.SelectUseCase(payload)
		if errUseCase != nil {
			return "error", errUseCase
		}

		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}

	return "messages processed", nil
}
