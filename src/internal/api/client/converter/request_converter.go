package converter

import (
	"github.com/amzdll/backend_2_go/src/internal/api/client/request"
	"github.com/amzdll/backend_2_go/src/internal/model"
)

func ToClientInfoFromRequest(clientReq request.CreationRequest) model.ClientInfo {
	return model.ClientInfo{
		ClientName:    clientReq.ClientName,
		ClientSurname: clientReq.ClientSurname,
		Birthday:      clientReq.Birthday,
		Gender:        clientReq.Gender,
		AddressId:     clientReq.AddressId,
	}
}

func ToClientFromUpdateAddressRequest(clientReq request.UpdationRequest) model.Client {
	return model.Client{
		Id:        clientReq.Id,
		AddressId: clientReq.AddressId,
	}
}
