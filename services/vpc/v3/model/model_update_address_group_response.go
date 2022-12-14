package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAddressGroupResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	AddressGroup   *AddressGroup `json:"address_group,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateAddressGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateAddressGroupResponse", string(data)}, " ")
}
