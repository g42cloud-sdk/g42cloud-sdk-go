package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListResourceGroupResponse struct {
	ResourceGroups *[]ResourceGroupInfo `json:"resource_groups,omitempty"`

	MetaData       *TotalMetaData `json:"meta_data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ListResourceGroupResponse", string(data)}, " ")
}
