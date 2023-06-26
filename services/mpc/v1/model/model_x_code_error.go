package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type XCodeError struct {
	Code *string `json:"code,omitempty"`

	Msg *string `json:"msg,omitempty"`
}

func (o XCodeError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "XCodeError struct{}"
	}

	return strings.Join([]string{"XCodeError", string(data)}, " ")
}
