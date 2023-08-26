package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ErrMsg struct {
	ErrorCode string `json:"error_code"`

	ErrorMsg string `json:"error_msg"`
}

func (o ErrMsg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrMsg struct{}"
	}

	return strings.Join([]string{"ErrMsg", string(data)}, " ")
}
