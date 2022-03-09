package useCase

import (
	"encoding/json"
	"github.com/melvi-co/zendesk-integation/internal/domain/model"
)

type UseCase interface {
	translate(payload json.RawMessage) (interface{}, error)
	Execute(payload model.Payload) error
}
