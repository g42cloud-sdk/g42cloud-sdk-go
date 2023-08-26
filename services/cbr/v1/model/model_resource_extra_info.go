package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResourceExtraInfo struct {
	ExcludeVolumes *[]string `json:"exclude_volumes,omitempty"`

	IncludeVolumes *[]ResourceExtraInfoIncludeVolumes `json:"include_volumes,omitempty"`
}

func (o ResourceExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtraInfo struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfo", string(data)}, " ")
}
