package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LvmConfig struct {
	LvType string `json:"lvType"`

	Path *string `json:"path,omitempty"`
}

func (o LvmConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LvmConfig struct{}"
	}

	return strings.Join([]string{"LvmConfig", string(data)}, " ")
}
