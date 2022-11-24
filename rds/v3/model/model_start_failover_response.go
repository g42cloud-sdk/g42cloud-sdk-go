package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartFailoverResponse struct {
	InstanceId *string `json:"instanceId,omitempty"`

	NodeId *string `json:"nodeId,omitempty"`

	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartFailoverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartFailoverResponse struct{}"
	}

	return strings.Join([]string{"StartFailoverResponse", string(data)}, " ")
}
