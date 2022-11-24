package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaSecurityGroupCommonRule struct {
	FromPort int32 `json:"from_port"`

	Group *NovaSecurityGroupCommonGroup `json:"group"`

	Id string `json:"id"`

	IpProtocol string `json:"ip_protocol"`

	IpRange *NovaSecurityGroupCommonIpRange `json:"ip_range"`

	ParentGroupId string `json:"parent_group_id"`

	ToPort int32 `json:"to_port"`
}

func (o NovaSecurityGroupCommonRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSecurityGroupCommonRule struct{}"
	}

	return strings.Join([]string{"NovaSecurityGroupCommonRule", string(data)}, " ")
}
