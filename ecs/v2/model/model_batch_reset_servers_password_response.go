package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResetServersPasswordResponse struct {
	Response       *[]ServerId `json:"response,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o BatchResetServersPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetServersPasswordResponse struct{}"
	}

	return strings.Join([]string{"BatchResetServersPasswordResponse", string(data)}, " ")
}
