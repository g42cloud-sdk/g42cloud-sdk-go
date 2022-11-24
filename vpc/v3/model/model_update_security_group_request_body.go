package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityGroupRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	SecurityGroup *UpdateSecurityGroupOption `json:"security_group"`
}

func (o UpdateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupRequestBody", string(data)}, " ")
}
