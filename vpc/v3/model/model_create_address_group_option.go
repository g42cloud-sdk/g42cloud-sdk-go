package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
