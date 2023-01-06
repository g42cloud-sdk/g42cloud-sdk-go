package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVolumeResponse struct {
	Volume         *VolumeDetail `json:"volume,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeResponse", string(data)}, " ")
}
