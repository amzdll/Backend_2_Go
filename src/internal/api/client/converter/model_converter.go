package converter

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/response"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func ModelToClientResponse(clientModel model.Client) response.ClientResponse {
	return response.ClientResponse{
		Id:               clientModel.Id,
		ClientName:       clientModel.ClientName,
		ClientSurname:    clientModel.ClientSurname,
		Birthday:         clientModel.Birthday,
		Gender:           clientModel.Gender,
		RegistrationDate: clientModel.RegistrationDate,
		AddressId:        clientModel.AddressId,
	}
}
