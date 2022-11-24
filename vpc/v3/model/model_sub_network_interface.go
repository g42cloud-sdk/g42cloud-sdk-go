package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubNetworkInterface struct {
	Id string `json:"id"`

	VirsubnetId string `json:"virsubnet_id"`

	PrivateIpAddress string `json:"private_ip_address"`

	Ipv6IpAddress string `json:"ipv6_ip_address"`

	MacAddress string `json:"mac_address"`

	ParentDeviceId string `json:"parent_device_id"`

	ParentId string `json:"parent_id"`

	Description string `json:"description"`

	VpcId string `json:"vpc_id"`

	VlanId int32 `json:"vlan_id"`

	SecurityGroups []string `json:"security_groups"`

	Tags []string `json:"tags"`

	ProjectId string `json:"project_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o SubNetworkInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubNetworkInterface struct{}"
	}

	return strings.Join([]string{"SubNetworkInterface", string(data)}, " ")
}
