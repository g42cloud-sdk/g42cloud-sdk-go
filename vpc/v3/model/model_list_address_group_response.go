package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAddressGroupResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	AddressGroups *[]AddressGroup `json:"address_groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"ListAddressGroupResponse", string(data)}, " ")
}
