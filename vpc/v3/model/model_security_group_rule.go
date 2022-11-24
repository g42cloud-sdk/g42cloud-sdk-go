package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityGroupRule struct {
	Id string `json:"id"`

	Description string `json:"description"`

	SecurityGroupId string `json:"security_group_id"`

	Direction string `json:"direction"`

	Protocol string `json:"protocol"`

	Ethertype string `json:"ethertype"`

	Multiport string `json:"multiport"`

	Action string `json:"action"`

	Priority int32 `json:"priority"`

	RemoteGroupId string `json:"remote_group_id"`

	RemoteIpPrefix string `json:"remote_ip_prefix"`

	RemoteAddressGroupId string `json:"remote_address_group_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	ProjectId string `json:"project_id"`
}

func (o SecurityGroupRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupRule struct{}"
	}

	return strings.Join([]string{"SecurityGroupRule", string(data)}, " ")
}
