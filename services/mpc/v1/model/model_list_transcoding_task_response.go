package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTranscodingTaskResponse struct {
	TaskArray *[]QueryTranscodingsTaskResponse `json:"task_array,omitempty"`

	IsTruncated *int32 `json:"is_truncated,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTranscodingTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodingTaskResponse struct{}"
	}

	return strings.Join([]string{"ListTranscodingTaskResponse", string(data)}, " ")
}
