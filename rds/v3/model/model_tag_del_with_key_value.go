package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagDelWithKeyValue struct {
	Key string `json:"key"`

	Value *string `json:"value,omitempty"`
}

func (o TagDelWithKeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagDelWithKeyValue struct{}"
	}

	return strings.Join([]string{"TagDelWithKeyValue", string(data)}, " ")
}
