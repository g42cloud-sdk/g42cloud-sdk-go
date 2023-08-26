package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResourceTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o CreateResourceTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceTagRequestBody", string(data)}, " ")
}
