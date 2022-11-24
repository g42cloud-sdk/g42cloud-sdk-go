package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronListFirewallGroupsResponse struct {
	FirewallGroups *[]NeutronFirewallGroup `json:"firewall_groups,omitempty"`

	FirewallGroupsLinks *[]NeutronPageLink `json:"firewall_groups_links,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o NeutronListFirewallGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFirewallGroupsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListFirewallGroupsResponse", string(data)}, " ")
}
