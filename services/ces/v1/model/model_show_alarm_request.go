package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowAlarmRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId string `json:"alarm_id"`
}

func (o ShowAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmRequest", string(data)}, " ")
}
