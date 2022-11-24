package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSecurityGroupRulesRequest struct {
	Marker *string `json:"marker,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

func (o ListSecurityGroupRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesRequest", string(data)}, " ")
}
