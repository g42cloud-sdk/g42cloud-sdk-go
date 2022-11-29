package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAlarmRequest struct {
	AlarmId string `json:"alarm_id"`
}

func (o DeleteAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRequest", string(data)}, " ")
}
