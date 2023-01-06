package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteAlarmRulesRequest struct {
	ContentType string `json:"Content-Type"`

	Body *BatchDeleteAlarmsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmRulesRequest", string(data)}, " ")
}
