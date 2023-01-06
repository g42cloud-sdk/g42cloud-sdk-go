package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAlarmRulesResponse struct {
	AlarmId        *string `json:"alarm_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmRulesResponse", string(data)}, " ")
}
