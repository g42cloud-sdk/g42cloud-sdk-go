package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o ShowSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityGroupRequest", string(data)}, " ")
}
