package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
