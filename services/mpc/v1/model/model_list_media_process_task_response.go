package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListMediaProcessTaskResponse struct {
	TaskArray *[]MediaProcessTaskInfo `json:"task_array,omitempty"`

	IsTruncated *int32 `json:"is_truncated,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMediaProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"ListMediaProcessTaskResponse", string(data)}, " ")
}
