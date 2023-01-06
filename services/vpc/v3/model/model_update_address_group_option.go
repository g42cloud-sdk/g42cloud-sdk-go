package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAddressGroupOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	IpSet *[]string `json:"ip_set,omitempty"`
}

func (o UpdateAddressGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateAddressGroupOption", string(data)}, " ")
}
