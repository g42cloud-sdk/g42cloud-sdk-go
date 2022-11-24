package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeFlavorObject struct {
	SpecCode string `json:"spec_code"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ResizeFlavorObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeFlavorObject struct{}"
	}

	return strings.Join([]string{"ResizeFlavorObject", string(data)}, " ")
}
