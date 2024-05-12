package model

import "github.com/google/uuid"

type Address struct {
	id      uuid.UUID
	country string
	city    string
	street  string
}
