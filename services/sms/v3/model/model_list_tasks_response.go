package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTasksResponse struct {
	Count *int32 `json:"count,omitempty"`

	Tasks          *[]TasksResponseBody `json:"tasks,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
