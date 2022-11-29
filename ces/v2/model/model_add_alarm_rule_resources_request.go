package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddAlarmRuleResourcesRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId string `json:"alarm_id"`

	Body *ResourcesReqV2 `json:"body,omitempty"`
}

func (o AddAlarmRuleResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAlarmRuleResourcesRequest struct{}"
	}

	return strings.Join([]string{"AddAlarmRuleResourcesRequest", string(data)}, " ")
}
