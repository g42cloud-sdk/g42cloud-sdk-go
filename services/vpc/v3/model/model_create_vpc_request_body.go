package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	Vpc *CreateVpcOption `json:"vpc"`
}

func (o CreateVpcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpcRequestBody", string(data)}, " ")
}
