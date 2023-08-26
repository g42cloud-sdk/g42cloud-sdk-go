package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSubNetworkInterfaceResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterface *SubNetworkInterface `json:"sub_network_interface,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o CreateSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"CreateSubNetworkInterfaceResponse", string(data)}, " ")
}
