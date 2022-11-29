package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpErrorInfo struct {
	Code *string `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`
}

func (o OpErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpErrorInfo struct{}"
	}

	return strings.Join([]string{"OpErrorInfo", string(data)}, " ")
}
