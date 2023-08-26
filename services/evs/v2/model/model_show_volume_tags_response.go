package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVolumeTagsResponse struct {
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVolumeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeTagsResponse", string(data)}, " ")
}
