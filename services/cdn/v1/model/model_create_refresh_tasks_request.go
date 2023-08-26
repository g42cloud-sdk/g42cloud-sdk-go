package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateRefreshTasksRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *RefreshTaskRequest `json:"body,omitempty"`
}

func (o CreateRefreshTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRefreshTasksRequest struct{}"
	}

	return strings.Join([]string{"CreateRefreshTasksRequest", string(data)}, " ")
}
