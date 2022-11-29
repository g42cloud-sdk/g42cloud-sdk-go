package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteAlarmsRequestBody struct {
	AlarmIds []string `json:"alarm_ids"`
}

func (o BatchDeleteAlarmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmsRequestBody", string(data)}, " ")
}
