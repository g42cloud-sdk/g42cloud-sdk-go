package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSubNetworkInterfaceResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterface *SubNetworkInterface `json:"sub_network_interface,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfaceResponse", string(data)}, " ")
}
