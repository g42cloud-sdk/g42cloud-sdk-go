package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListVolumesByTagsRequest struct {
	Body *ListVolumesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListVolumesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListVolumesByTagsRequest", string(data)}, " ")
}
