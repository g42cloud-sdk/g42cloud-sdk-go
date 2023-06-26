package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMergeChannelsTaskResponse struct {
	TaskArray *[]MergeChannelsTaskInfo `json:"task_array,omitempty"`

	IsTruncated *int32 `json:"is_truncated,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMergeChannelsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeChannelsTaskResponse struct{}"
	}

	return strings.Join([]string{"ListMergeChannelsTaskResponse", string(data)}, " ")
}
