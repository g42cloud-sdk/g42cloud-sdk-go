package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePreheatingTasksResponse struct {
	PreheatingTask *string `json:"preheating_task,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePreheatingTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePreheatingTasksResponse struct{}"
	}

	return strings.Join([]string{"CreatePreheatingTasksResponse", string(data)}, " ")
}
