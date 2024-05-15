package request

import "github.com/google/uuid"

type UpdationRequest struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	AddressId uuid.UUID `json:"address_id" validate:"required"`
}
