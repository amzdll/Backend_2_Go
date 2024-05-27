package response

import (
	"github.com/amzdll/shop/internal/model"
	"github.com/google/uuid"
	"time"
)

type ClientResponse struct {
	Id               uuid.UUID `json:"id"`
	ClientName       string    `json:"client_name"`
	ClientSurname    string    `json:"client_surname"`
	Birthday         time.Time `json:"birthday"`
	Gender           string    `json:"gender"`
	RegistrationDate time.Time `json:"registration_date"`
	AddressId        uuid.UUID `json:"address_id"`
}

func (cr *ClientResponse) From(m model.Client) {
	cr.Id = m.Id
	cr.ClientName = m.ClientName
	cr.ClientSurname = m.ClientSurname
	cr.Birthday = m.Birthday
	cr.Gender = m.Gender
	cr.RegistrationDate = m.RegistrationDate
	cr.AddressId = m.AddressId
}
