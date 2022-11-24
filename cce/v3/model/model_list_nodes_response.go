package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListNodesResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Items          *[]Node `json:"items,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodesResponse struct{}"
	}

	return strings.Join([]string{"ListNodesResponse", string(data)}, " ")
}
