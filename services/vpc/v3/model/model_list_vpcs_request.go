package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListVpcsRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	Description *[]string `json:"description,omitempty"`

	Cidr *[]string `json:"cidr,omitempty"`
}

func (o ListVpcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcsRequest", string(data)}, " ")
}
