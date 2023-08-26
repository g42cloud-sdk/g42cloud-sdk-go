package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
