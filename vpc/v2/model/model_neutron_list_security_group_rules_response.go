package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronListSecurityGroupRulesResponse struct {
	SecurityGroupRules *[]NeutronSecurityGroupRule `json:"security_group_rules,omitempty"`

	SecurityGroupRulesLinks *[]NeutronPageLink `json:"security_group_rules_links,omitempty"`
	HttpStatusCode          int                `json:"-"`
}

func (o NeutronListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupRulesResponse", string(data)}, " ")
}
