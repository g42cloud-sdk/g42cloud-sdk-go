package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSecurityGroupsResponse struct {
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecurityGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsResponse", string(data)}, " ")
}
