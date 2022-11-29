package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlarmHistoriesRequest struct {
	ContentType string `json:"Content-Type"`

	AlarmId *string `json:"alarm_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	Level *int32 `json:"level,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	ResourceId *string `json:"resource_id,omitempty"`

	From *string `json:"from,omitempty"`

	To *string `json:"to,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}
