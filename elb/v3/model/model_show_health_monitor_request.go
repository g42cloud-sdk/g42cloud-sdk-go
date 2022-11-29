package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHealthMonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o ShowHealthMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthMonitorRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthMonitorRequest", string(data)}, " ")
}
