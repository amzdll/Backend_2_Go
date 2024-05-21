package request

import (
	"github.com/amzdll/backend_2_go/src/internal/model"
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

func (cr *CreateRequest) ToClientInfo() model.ClientInfo {
	return model.ClientInfo{
		ClientName:    cr.ClientName,
		ClientSurname: cr.ClientSurname,
		Birthday:      cr.Birthday,
		Gender:        cr.Gender,
		AddressId:     cr.AddressId,
	}
}
