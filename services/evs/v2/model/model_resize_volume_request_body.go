package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResizeVolumeRequestBody struct {
	BssParam *BssParamForResizeVolume `json:"bssParam,omitempty"`

	OsExtend *OsExtend `json:"os-extend"`
}

func (o ResizeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequestBody", string(data)}, " ")
}
