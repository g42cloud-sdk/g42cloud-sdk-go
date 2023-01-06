package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAddressGroupRequest struct {
	AddressGroupId string `json:"address_group_id"`

	Body *UpdateAddressGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateAddressGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateAddressGroupRequest", string(data)}, " ")
}
