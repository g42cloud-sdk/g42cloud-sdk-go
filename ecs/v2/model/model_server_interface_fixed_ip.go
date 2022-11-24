package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerInterfaceFixedIp struct {
	IpAddress *string `json:"ip_address,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o ServerInterfaceFixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerInterfaceFixedIp struct{}"
	}

	return strings.Join([]string{"ServerInterfaceFixedIp", string(data)}, " ")
}
