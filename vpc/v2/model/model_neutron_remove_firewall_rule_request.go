package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
