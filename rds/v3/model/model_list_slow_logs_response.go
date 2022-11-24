package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSlowLogsResponse struct {
	SlowLogList *[]SlowLog `json:"slow_log_list,omitempty"`

	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogsResponse", string(data)}, " ")
}
