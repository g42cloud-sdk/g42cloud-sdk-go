package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegisterServerMonitorRequest struct {
	ServerId string `json:"server_id"`

	Body *RegisterServerMonitorRequestBody `json:"body,omitempty"`
}

func (o RegisterServerMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerMonitorRequest struct{}"
	}

	return strings.Join([]string{"RegisterServerMonitorRequest", string(data)}, " ")
}
