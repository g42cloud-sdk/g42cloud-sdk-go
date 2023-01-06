package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteAlarmRulesResponse struct {
	AlarmIds       *[]string `json:"alarm_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmRulesResponse", string(data)}, " ")
}
