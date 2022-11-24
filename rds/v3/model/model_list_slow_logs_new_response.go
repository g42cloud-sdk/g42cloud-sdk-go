package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSlowLogsNewResponse struct {
	SlowLogList *[]SlowLog `json:"slow_log_list,omitempty"`

	TotalRecord    *int32 `json:"total_record,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogsNewResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogsNewResponse", string(data)}, " ")
}
