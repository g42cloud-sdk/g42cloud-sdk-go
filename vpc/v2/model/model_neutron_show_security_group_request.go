package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronShowSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o NeutronShowSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupRequest", string(data)}, " ")
}
