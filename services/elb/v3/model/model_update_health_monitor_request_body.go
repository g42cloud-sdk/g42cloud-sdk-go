package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateHealthMonitorRequestBody struct {
	Healthmonitor *UpdateHealthMonitorOption `json:"healthmonitor"`
}

func (o UpdateHealthMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorRequestBody", string(data)}, " ")
}
