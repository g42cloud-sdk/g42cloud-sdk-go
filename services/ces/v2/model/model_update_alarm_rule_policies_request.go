package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAlarmRulePoliciesRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId string `json:"alarm_id"`

	Body *PoliciesReqV2 `json:"body,omitempty"`
}

func (o UpdateAlarmRulePoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRulePoliciesRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRulePoliciesRequest", string(data)}, " ")
}
