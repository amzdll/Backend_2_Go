package request

import "github.com/google/uuid"

type DeleteClientRequest struct {
	Id uuid.UUID `json:"id" validate:"required"`
}
