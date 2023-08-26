package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SetSecurityGroupResponse struct {
	WorkflowId     *string `json:"workflowId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"SetSecurityGroupResponse", string(data)}, " ")
}
