package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`

	CurrentCount int32 `json:"current_count"`

	NextMarker *string `json:"next_marker,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
