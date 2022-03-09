package selector

import (
	"fmt"
	"github.com/melvi-co/zendesk-integation/internal/domain/model"
	"github.com/melvi-co/zendesk-integation/internal/domain/useCase"
)

type Selector interface {
	SelectUseCase(payload model.Payload) error
}

type useCaseSelector struct {
	createUserAndOrganizationUC useCase.UseCase
}

func NewUseCase() Selector {
	return &useCaseSelector{
		createUserAndOrganizationUC: useCase.NewCreateUserAndOrganizationUC(),
	}
}

func (uc useCaseSelector) SelectUseCase(payload model.Payload) error {
	event := payload.Event
	switch event {
	case model.CreateUserAndOrganization:
		return uc.createUserAndOrganizationUC.Execute(payload)
	default:
		return fmt.Errorf("useCase not found: %s \n Payload %v", event, payload.Payload)
	}
}
