package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveVpcExtendCidrRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	Vpc *RemoveExtendCidrOption `json:"vpc"`
}

func (o RemoveVpcExtendCidrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVpcExtendCidrRequestBody struct{}"
	}

	return strings.Join([]string{"RemoveVpcExtendCidrRequestBody", string(data)}, " ")
}
