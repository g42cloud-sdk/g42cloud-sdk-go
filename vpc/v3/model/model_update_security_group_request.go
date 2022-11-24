package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`

	Body *UpdateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupRequest", string(data)}, " ")
}
