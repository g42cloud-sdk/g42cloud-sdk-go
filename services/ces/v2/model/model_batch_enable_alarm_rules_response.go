package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchEnableAlarmRulesResponse struct {
	AlarmIds       *[]string `json:"alarm_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchEnableAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmRulesResponse", string(data)}, " ")
}
