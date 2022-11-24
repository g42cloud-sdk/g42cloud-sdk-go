package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVolumeRequest struct {
	VolumeId string `json:"volume_id"`

	Body *UpdateVolumeRequestBody `json:"body,omitempty"`
}

func (o UpdateVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeRequest struct{}"
	}

	return strings.Join([]string{"UpdateVolumeRequest", string(data)}, " ")
}
