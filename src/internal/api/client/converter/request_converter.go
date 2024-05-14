package converter

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/request"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func ToClientInfoFromRequest(clientReq request.CreateClientRequest) model.ClientInfo {
	return model.ClientInfo{
		ClientName:    clientReq.ClientName,
		ClientSurname: clientReq.ClientSurname,
		Birthday:      clientReq.Birthday,
		Gender:        clientReq.Gender,
		AddressId:     clientReq.AddressId,
	}
}

func ToClientListParamsFromRequest(clientReq request.GetAllRequest) model.Pagination {
	return model.Pagination{
		Limit:  clientReq.Limit,
		Offset: clientReq.Offset,
	}
}
