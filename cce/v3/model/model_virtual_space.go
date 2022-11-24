package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VirtualSpace struct {
	Name string `json:"name"`

	Size string `json:"size"`

	LvmConfig *LvmConfig `json:"lvmConfig,omitempty"`

	RuntimeConfig *RuntimeConfig `json:"runtimeConfig,omitempty"`
}

func (o VirtualSpace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualSpace struct{}"
	}

	return strings.Join([]string{"VirtualSpace", string(data)}, " ")
}
