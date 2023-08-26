package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMergeChannelsTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMergeChannelsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeChannelsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateMergeChannelsTaskResponse", string(data)}, " ")
}
