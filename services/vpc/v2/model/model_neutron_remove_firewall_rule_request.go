package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronRemoveFirewallRuleRequest struct {
	FirewallPolicyId string `json:"firewall_policy_id"`

	Body *NeutronRemoveFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronRemoveFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRemoveFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronRemoveFirewallRuleRequest", string(data)}, " ")
}
