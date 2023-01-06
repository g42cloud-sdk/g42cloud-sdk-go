package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ModifyAlarmActionReq struct {
	AlarmEnabled bool `json:"alarm_enabled"`
}

func (o ModifyAlarmActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAlarmActionReq struct{}"
	}

	return strings.Join([]string{"ModifyAlarmActionReq", string(data)}, " ")
}
