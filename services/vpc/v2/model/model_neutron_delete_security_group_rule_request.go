package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteSecurityGroupRuleRequest struct {
	SecurityGroupRuleId string `json:"security_group_rule_id"`
}

func (o NeutronDeleteSecurityGroupRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupRuleRequest", string(data)}, " ")
}
