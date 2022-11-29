package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAlarmRulesRequest struct {
	ContentType string `json:"Content-Type"`

	Body *PostAlarmsReqV2 `json:"body,omitempty"`
}

func (o CreateAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmRulesRequest", string(data)}, " ")
}
