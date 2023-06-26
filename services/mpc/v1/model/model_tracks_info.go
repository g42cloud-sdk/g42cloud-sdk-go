package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TracksInfo struct {
	ChannelLayout *string `json:"channel_layout,omitempty"`

	Language *string `json:"language,omitempty"`
}

func (o TracksInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TracksInfo struct{}"
	}

	return strings.Join([]string{"TracksInfo", string(data)}, " ")
}
