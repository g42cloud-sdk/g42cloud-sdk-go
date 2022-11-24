package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateServersRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	Server *PrePaidServer `json:"server"`
}

func (o CreateServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServersRequestBody struct{}"
	}

	return strings.Join([]string{"CreateServersRequestBody", string(data)}, " ")
}
