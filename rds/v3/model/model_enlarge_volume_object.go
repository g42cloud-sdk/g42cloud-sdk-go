package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnlargeVolumeObject struct {
	Size int32 `json:"size"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o EnlargeVolumeObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeVolumeObject struct{}"
	}

	return strings.Join([]string{"EnlargeVolumeObject", string(data)}, " ")
}
