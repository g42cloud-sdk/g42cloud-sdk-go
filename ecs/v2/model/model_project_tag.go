package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectTag struct {
	Key string `json:"key"`

	Values *[]string `json:"values,omitempty"`
}

func (o ProjectTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTag struct{}"
	}

	return strings.Join([]string{"ProjectTag", string(data)}, " ")
}
