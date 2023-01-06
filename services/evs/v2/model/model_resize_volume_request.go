package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResizeVolumeRequest struct {
	VolumeId string `json:"volume_id"`

	Body *ResizeVolumeRequestBody `json:"body,omitempty"`
}

func (o ResizeVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequest struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequest", string(data)}, " ")
}
