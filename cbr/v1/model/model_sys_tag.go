package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SysTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o SysTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTag struct{}"
	}

	return strings.Join([]string{"SysTag", string(data)}, " ")
}
