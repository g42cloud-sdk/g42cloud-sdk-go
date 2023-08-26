package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddVpcExtendCidrRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	Vpc *AddExtendCidrOption `json:"vpc"`
}

func (o AddVpcExtendCidrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVpcExtendCidrRequestBody struct{}"
	}

	return strings.Join([]string{"AddVpcExtendCidrRequestBody", string(data)}, " ")
}
