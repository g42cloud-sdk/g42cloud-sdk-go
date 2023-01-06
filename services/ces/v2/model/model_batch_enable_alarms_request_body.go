package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
