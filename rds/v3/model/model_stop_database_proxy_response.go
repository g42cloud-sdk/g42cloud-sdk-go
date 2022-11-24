package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StopDatabaseProxyResponse struct {
	WorkflowId     *string `json:"workflow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopDatabaseProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopDatabaseProxyResponse struct{}"
	}

	return strings.Join([]string{"StopDatabaseProxyResponse", string(data)}, " ")
}
