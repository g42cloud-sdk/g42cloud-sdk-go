package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateSubNetworkInterfaceResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterfaces *[]SubNetworkInterface `json:"sub_network_interfaces,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o BatchCreateSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSubNetworkInterfaceResponse", string(data)}, " ")
}
