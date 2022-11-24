package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronDeleteFirewallGroupRequest struct {
	FirewallGroupId string `json:"firewall_group_id"`
}

func (o NeutronDeleteFirewallGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallGroupRequest", string(data)}, " ")
}
