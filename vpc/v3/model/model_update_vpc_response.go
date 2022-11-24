package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpcResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpcResponse", string(data)}, " ")
}
