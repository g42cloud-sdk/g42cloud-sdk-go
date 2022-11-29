package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchEnableAlarmsRequestBody struct {
	AlarmIds []string `json:"alarm_ids"`

	AlarmEnabled bool `json:"alarm_enabled"`
}

func (o BatchEnableAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmsRequestBody", string(data)}, " ")
}
