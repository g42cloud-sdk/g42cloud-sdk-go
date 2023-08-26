package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type StartDatabaseProxyResponse struct {
	WorkflowId     *string `json:"workflow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartDatabaseProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDatabaseProxyResponse struct{}"
	}

	return strings.Join([]string{"StartDatabaseProxyResponse", string(data)}, " ")
}
