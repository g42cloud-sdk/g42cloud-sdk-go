package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAlarmHistoriesResponse struct {
	AlarmHistories *[]AlarmHistoryItemV2 `json:"alarm_histories,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesResponse", string(data)}, " ")
}
