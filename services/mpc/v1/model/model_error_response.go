package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ErrorResponse struct {
	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ErrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorResponse struct{}"
	}

	return strings.Join([]string{"ErrorResponse", string(data)}, " ")
}
