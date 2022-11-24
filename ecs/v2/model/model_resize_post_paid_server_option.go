package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizePostPaidServerOption struct {
	FlavorRef string `json:"flavorRef"`

	Mode *string `json:"mode,omitempty"`
}

func (o ResizePostPaidServerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerOption struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerOption", string(data)}, " ")
}
