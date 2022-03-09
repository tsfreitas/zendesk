package useCase

import (
	"encoding/json"
	"fmt"
	"github.com/melvi-co/zendesk-integation/internal/domain/model"
)

type createUseAndOrganizationUC struct {
	zendeskAdapter string
}

func NewCreateUserAndOrganizationUC() UseCase {
	return &createUseAndOrganizationUC{
		zendeskAdapter: "",
	}
}

func (c createUseAndOrganizationUC) translate(payload json.RawMessage) (interface{}, error) {
	var model = new(model.CreateUserAndOrganizationPayload)
	err := json.Unmarshal(payload, &model)
	return model, err
}

func (c createUseAndOrganizationUC) Execute(payload model.Payload) error {
	translated, err := c.translate(payload.Payload)
	event := translated.(*model.CreateUserAndOrganizationPayload)

	// Cria usuário no zendesk
	// Cria organization
	// Associa o usuário à organization

	fmt.Printf("Print my event payload %v \n", event)
	return err
}
