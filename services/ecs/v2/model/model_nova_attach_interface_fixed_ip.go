package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaAttachInterfaceFixedIp struct {
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o NovaAttachInterfaceFixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAttachInterfaceFixedIp struct{}"
	}

	return strings.Join([]string{"NovaAttachInterfaceFixedIp", string(data)}, " ")
}
