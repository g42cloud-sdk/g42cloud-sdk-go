package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateFollowerResponse struct {
	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateFollowerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateFollowerResponse struct{}"
	}

	return strings.Join([]string{"MigrateFollowerResponse", string(data)}, " ")
}
