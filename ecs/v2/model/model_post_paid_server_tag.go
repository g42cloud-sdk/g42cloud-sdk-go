package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostPaidServerTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o PostPaidServerTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerTag struct{}"
	}

	return strings.Join([]string{"PostPaidServerTag", string(data)}, " ")
}
