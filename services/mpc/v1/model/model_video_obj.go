package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VideoObj struct {
	Width int32 `json:"width"`

	Height int32 `json:"height"`

	Bitrate int32 `json:"bitrate"`
}

func (o VideoObj) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoObj struct{}"
	}

	return strings.Join([]string{"VideoObj", string(data)}, " ")
}
