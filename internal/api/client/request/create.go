package request

import (
	"github.com/amzdll/shop/internal/model"
	"github.com/google/uuid"
	"time"
)

type CreateRequest struct {
	ClientName    string    `json:"name" validate:"required"`
	ClientSurname string    `json:"surname" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required"`
	Gender        string    `json:"gender" validate:"required"`
	AddressId     uuid.UUID `json:"address_id" validate:"required"`
}

func (r *CreateRequest) ToClientInfo() model.ClientInfo {
	return model.ClientInfo{
		ClientName:    r.ClientName,
		ClientSurname: r.ClientSurname,
		Birthday:      r.Birthday,
		Gender:        r.Gender,
		AddressId:     r.AddressId,
	}
}
