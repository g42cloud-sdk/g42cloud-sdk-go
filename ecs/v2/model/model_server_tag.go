package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o ServerTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerTag struct{}"
	}

	return strings.Join([]string{"ServerTag", string(data)}, " ")
}
