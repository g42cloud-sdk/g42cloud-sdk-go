package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateHealthMonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`

	Body *UpdateHealthMonitorRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorRequest", string(data)}, " ")
}
