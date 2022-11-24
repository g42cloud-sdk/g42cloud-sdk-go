package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSecurityGroupRuleOption struct {
	SecurityGroupId string `json:"security_group_id"`

	Description *string `json:"description,omitempty"`

	Direction string `json:"direction"`

	Ethertype *string `json:"ethertype,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	Multiport *string `json:"multiport,omitempty"`

	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`

	RemoteGroupId *string `json:"remote_group_id,omitempty"`

	RemoteAddressGroupId *string `json:"remote_address_group_id,omitempty"`

	Action *string `json:"action,omitempty"`

	Priority *string `json:"priority,omitempty"`
}

func (o CreateSecurityGroupRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleOption", string(data)}, " ")
}
