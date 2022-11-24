package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaSecurityGroup struct {
	Description string `json:"description"`

	Id string `json:"id"`

	Name string `json:"name"`

	TenantId string `json:"tenant_id"`

	Rules []NovaSecurityGroupCommonRule `json:"rules"`
}

func (o NovaSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSecurityGroup struct{}"
	}

	return strings.Join([]string{"NovaSecurityGroup", string(data)}, " ")
}
