package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronShowFirewallGroupRequest struct {
	FirewallGroupId string `json:"firewall_group_id"`
}

func (o NeutronShowFirewallGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallGroupRequest", string(data)}, " ")
}
