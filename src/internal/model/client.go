package model

import (
	"github.com/google/uuid"
	"time"
)

type Client struct {
	Id               uuid.UUID
	ClientName       string
	ClientSurname    string
	Birthday         time.Time
	Gender           string
	RegistrationDate time.Time
	AddressId        uuid.UUID
}

type ClientInfo struct {
	ClientName       string
	ClientSurname    string
	Birthday         time.Time
	Gender           string
	RegistrationDate time.Time
	AddressId        uuid.UUID
}

type ClientListParams struct {
	Limit  uint64
	Offset uint64
}
