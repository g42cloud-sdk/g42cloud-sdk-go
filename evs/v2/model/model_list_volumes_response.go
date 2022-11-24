package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVolumesResponse struct {
	Count *int32 `json:"count,omitempty"`

	VolumesLinks *[]Link `json:"volumes_links,omitempty"`

	Volumes        *[]VolumeDetail `json:"volumes,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesResponse", string(data)}, " ")
}
