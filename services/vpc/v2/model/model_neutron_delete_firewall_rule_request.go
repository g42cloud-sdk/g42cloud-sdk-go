package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
