package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAlarmHistoriesResponse struct {
	AlarmHistories *[]AlarmHistoryInfo `json:"alarm_histories,omitempty"`

	MetaData       *MetaDataForAlarmHistory `json:"meta_data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListAlarmHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesResponse", string(data)}, " ")
}
