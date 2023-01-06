package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowAddressGroupResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	AddressGroup   *AddressGroup `json:"address_group,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowAddressGroupResponse", string(data)}, " ")
}
