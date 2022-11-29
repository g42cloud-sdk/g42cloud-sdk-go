package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
