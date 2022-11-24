package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePreheatingTasksRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *PreheatingTaskRequest `json:"body,omitempty"`
}

func (o CreatePreheatingTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePreheatingTasksRequest struct{}"
	}

	return strings.Join([]string{"CreatePreheatingTasksRequest", string(data)}, " ")
}
