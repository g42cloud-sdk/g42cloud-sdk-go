package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CinderListVolumeTypesResponse struct {
	VolumeTypes    *[]VolumeType `json:"volume_types,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CinderListVolumeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTypesResponse struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTypesResponse", string(data)}, " ")
}
