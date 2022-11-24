package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronShowFirewallPolicyRequest struct {
	FirewallPolicyId string `json:"firewall_policy_id"`
}

func (o NeutronShowFirewallPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallPolicyRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallPolicyRequest", string(data)}, " ")
}
