package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVolumeTagsResponse struct {
	Tags           map[string][]string `json:"tags,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVolumeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"ListVolumeTagsResponse", string(data)}, " ")
}
