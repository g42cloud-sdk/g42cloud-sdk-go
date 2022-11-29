package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlarmActionRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId string `json:"alarm_id"`

	Body *ModifyAlarmActionReq `json:"body,omitempty"`
}

func (o UpdateAlarmActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionRequest", string(data)}, " ")
}
