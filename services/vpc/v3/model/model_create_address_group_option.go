package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAddressGroupOption struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	IpVersion int32 `json:"ip_version"`

	IpSet *[]string `json:"ip_set,omitempty"`
}

func (o CreateAddressGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddressGroupOption struct{}"
	}

	return strings.Join([]string{"CreateAddressGroupOption", string(data)}, " ")
}
