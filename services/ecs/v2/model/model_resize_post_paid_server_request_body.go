package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResizePostPaidServerRequestBody struct {
	Resize *ResizePostPaidServerOption `json:"resize"`

	DryRun *bool `json:"dry_run,omitempty"`
}

func (o ResizePostPaidServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerRequestBody struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerRequestBody", string(data)}, " ")
}
