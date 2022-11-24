package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListNodePoolsResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Items          *[]NodePool `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListNodePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodePoolsResponse struct{}"
	}

	return strings.Join([]string{"ListNodePoolsResponse", string(data)}, " ")
}
