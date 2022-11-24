package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerId struct {
	Id string `json:"id"`
}

func (o ServerId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerId struct{}"
	}

	return strings.Join([]string{"ServerId", string(data)}, " ")
}
