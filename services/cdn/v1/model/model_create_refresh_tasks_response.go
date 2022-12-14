package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateRefreshTasksResponse struct {
	RefreshTask    *string `json:"refresh_task,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRefreshTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRefreshTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateRefreshTasksResponse", string(data)}, " ")
}
