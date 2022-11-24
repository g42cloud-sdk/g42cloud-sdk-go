package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAddressGroupRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	AddressGroup *CreateAddressGroupOption `json:"address_group"`
}

func (o CreateAddressGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddressGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAddressGroupRequestBody", string(data)}, " ")
}
