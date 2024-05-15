package request

import (
	"github.com/google/uuid"
	"time"
)

type CreationRequest struct {
	ClientName    string    `json:"name" validate:"required"`
	ClientSurname string    `json:"surname" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required"`
	Gender        string    `json:"gender" validate:"required"`
	AddressId     uuid.UUID `json:"address_id" validate:"required"`
}
