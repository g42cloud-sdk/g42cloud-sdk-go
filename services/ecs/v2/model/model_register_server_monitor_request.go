package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
