package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpcResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcResponse", string(data)}, " ")
}
