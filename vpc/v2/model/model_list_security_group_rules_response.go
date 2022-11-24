package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSecurityGroupRulesResponse struct {
	SecurityGroupRules *[]SecurityGroupRule `json:"security_group_rules,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesResponse", string(data)}, " ")
}
