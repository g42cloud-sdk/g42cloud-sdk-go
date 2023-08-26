package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o NeutronDeleteSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupRequest", string(data)}, " ")
}
