package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
