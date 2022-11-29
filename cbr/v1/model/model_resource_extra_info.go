package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceExtraInfo struct {
	ExcludeVolumes *[]string `json:"exclude_volumes,omitempty"`
}

func (o ResourceExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtraInfo struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfo", string(data)}, " ")
}
