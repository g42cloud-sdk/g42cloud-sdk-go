package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetServerPasswordRequest struct {
	ServerId string `json:"server_id"`

	Body *ResetServerPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetServerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordRequest", string(data)}, " ")
}
