package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListVpcsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Vpcs *[]Vpc `json:"vpcs,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcsResponse", string(data)}, " ")
}
