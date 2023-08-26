package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o DeleteSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupRequest", string(data)}, " ")
}
