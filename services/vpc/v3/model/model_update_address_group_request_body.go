package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAddressGroupRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	AddressGroup *UpdateAddressGroupOption `json:"address_group"`
}

func (o UpdateAddressGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAddressGroupRequestBody", string(data)}, " ")
}
