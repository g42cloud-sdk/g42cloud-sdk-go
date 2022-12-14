package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSubNetworkInterfaceRequest struct {
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`

	Body *UpdateSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o UpdateSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubNetworkInterfaceRequest", string(data)}, " ")
}
