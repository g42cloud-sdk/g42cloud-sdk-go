package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceBody struct {
	Sources []SourceWithPort `json:"sources"`
}

func (o ResourceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceBody struct{}"
	}

	return strings.Join([]string{"ResourceBody", string(data)}, " ")
}
