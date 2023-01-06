package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAlarmResponse struct {
	AlarmId        *string `json:"alarm_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmResponse", string(data)}, " ")
}
