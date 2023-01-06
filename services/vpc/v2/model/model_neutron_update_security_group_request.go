package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`

	Body *NeutronUpdateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupRequest", string(data)}, " ")
}
