package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMergeChannelsTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteMergeChannelsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMergeChannelsTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMergeChannelsTaskRequest", string(data)}, " ")
}
