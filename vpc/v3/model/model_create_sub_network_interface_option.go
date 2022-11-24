package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubNetworkInterfaceOption struct {
	VirsubnetId string `json:"virsubnet_id"`

	VlanId *string `json:"vlan_id,omitempty"`

	ParentId string `json:"parent_id"`

	Description *string `json:"description,omitempty"`

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	Ipv6IpAddress *string `json:"ipv6_ip_address,omitempty"`

	SecurityGroups *[]string `json:"security_groups,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`
}

func (o CreateSubNetworkInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubNetworkInterfaceOption struct{}"
	}

	return strings.Join([]string{"CreateSubNetworkInterfaceOption", string(data)}, " ")
}
