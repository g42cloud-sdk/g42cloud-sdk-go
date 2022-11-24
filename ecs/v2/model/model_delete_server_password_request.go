package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteServerPasswordRequest struct {
	ServerId string `json:"server_id"`
}

func (o DeleteServerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerPasswordRequest", string(data)}, " ")
}
