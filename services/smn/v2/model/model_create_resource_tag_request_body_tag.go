package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResourceTagRequestBodyTag struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o CreateResourceTagRequestBodyTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagRequestBodyTag struct{}"
	}

	return strings.Join([]string{"CreateResourceTagRequestBodyTag", string(data)}, " ")
}
