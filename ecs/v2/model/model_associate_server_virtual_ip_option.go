package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateServerVirtualIpOption struct {
	SubnetId string `json:"subnet_id"`

	IpAddress string `json:"ip_address"`

	ReverseBinding *bool `json:"reverse_binding,omitempty"`
}

func (o AssociateServerVirtualIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateServerVirtualIpOption struct{}"
	}

	return strings.Join([]string{"AssociateServerVirtualIpOption", string(data)}, " ")
}
