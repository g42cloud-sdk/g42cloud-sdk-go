package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResourceExtraInfoIncludeVolumes struct {
	Id string `json:"id"`

	OsVersion *string `json:"os_version,omitempty"`
}

func (o ResourceExtraInfoIncludeVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtraInfoIncludeVolumes struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfoIncludeVolumes", string(data)}, " ")
}
