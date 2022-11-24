package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateServerRequest struct {
	ServerId string `json:"server_id"`

	Body *UpdateServerRequestBody `json:"body,omitempty"`
}

func (o UpdateServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerRequest", string(data)}, " ")
}
