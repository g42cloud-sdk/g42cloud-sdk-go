package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AudioTrack struct {
	Type *int32 `json:"type,omitempty"`

	Left *int32 `json:"left,omitempty"`

	Right *int32 `json:"right,omitempty"`
}

func (o AudioTrack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioTrack struct{}"
	}

	return strings.Join([]string{"AudioTrack", string(data)}, " ")
}
