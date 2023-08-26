package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSubNetworkInterfacesQuantityResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterfaces *int32 `json:"sub_network_interfaces,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ShowSubNetworkInterfacesQuantityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfacesQuantityResponse struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfacesQuantityResponse", string(data)}, " ")
}
