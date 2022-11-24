package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHistoryTasksResponse struct {
	Total *int32 `json:"total,omitempty"`

	Tasks          *[]TasksObject `json:"tasks,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowHistoryTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTasksResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTasksResponse", string(data)}, " ")
}
