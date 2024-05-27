package request

import (
	"github.com/amzdll/shop/internal/model"
	"github.com/google/uuid"
)

type UpdateRequest struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	AddressId uuid.UUID `json:"address_id" validate:"required"`
}

func (r *UpdateRequest) ToClient() model.Client {
	return model.Client{
		Id:        r.Id,
		AddressId: r.AddressId,
	}
}
