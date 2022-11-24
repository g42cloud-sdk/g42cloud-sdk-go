package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpResponseHeader struct {
	Name string `json:"name"`

	Value *string `json:"value,omitempty"`

	Action string `json:"action"`
}

func (o HttpResponseHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpResponseHeader struct{}"
	}

	return strings.Join([]string{"HttpResponseHeader", string(data)}, " ")
}
