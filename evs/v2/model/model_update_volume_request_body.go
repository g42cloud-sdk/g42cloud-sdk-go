package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVolumeRequestBody struct {
	Volume *UpdateVolumeOption `json:"volume"`
}

func (o UpdateVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVolumeRequestBody", string(data)}, " ")
}
