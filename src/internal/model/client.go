package model

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	Id               uuid.UUID `db:"id"`
	ClientName       string    `db:"client_name"`
	ClientSurname    string    `db:"client_surname"`
	Birthday         time.Time `db:"birthday"`
	Gender           string    `db:"gender"`
	RegistrationDate time.Time `db:"registration_date"`
	AddressId        uuid.UUID `db:"address_id"`
}

type ClientInfo struct {
	ClientName    string    `db:"client_name"`
	ClientSurname string    `db:"client_surname"`
	Birthday      time.Time `db:"birthday"`
	Gender        string    `db:"gender"`
	AddressId     uuid.UUID `db:"address_id"`
}
