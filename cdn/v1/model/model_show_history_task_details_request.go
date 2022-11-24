package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHistoryTaskDetailsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	HistoryTasksId string `json:"history_tasks_id"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNumber *int32 `json:"page_number,omitempty"`

	Status *string `json:"status,omitempty"`

	Url *string `json:"url,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ShowHistoryTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryTaskDetailsRequest", string(data)}, " ")
}
