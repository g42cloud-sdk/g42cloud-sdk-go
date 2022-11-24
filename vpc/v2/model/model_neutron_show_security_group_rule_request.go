package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronShowSecurityGroupRuleRequest struct {
	SecurityGroupRuleId string `json:"security_group_rule_id"`
}

func (o NeutronShowSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupRuleRequest", string(data)}, " ")
}
