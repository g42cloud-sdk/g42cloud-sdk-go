package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronDeleteFirewallRuleRequest struct {
	FirewallRuleId string `json:"firewall_rule_id"`
}

func (o NeutronDeleteFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallRuleRequest", string(data)}, " ")
}
