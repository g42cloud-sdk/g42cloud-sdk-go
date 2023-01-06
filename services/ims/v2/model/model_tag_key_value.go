package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TagKeyValue struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o TagKeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagKeyValue struct{}"
	}

	return strings.Join([]string{"TagKeyValue", string(data)}, " ")
}
