package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSecurityGroupRulesRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *[]string `json:"id,omitempty"`

	SecurityGroupId *[]string `json:"security_group_id,omitempty"`

	Protocol *[]string `json:"protocol,omitempty"`

	Description *[]string `json:"description,omitempty"`

	RemoteGroupId *[]string `json:"remote_group_id,omitempty"`

	Direction *string `json:"direction,omitempty"`

	Action *string `json:"action,omitempty"`
}

func (o ListSecurityGroupRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesRequest", string(data)}, " ")
}
