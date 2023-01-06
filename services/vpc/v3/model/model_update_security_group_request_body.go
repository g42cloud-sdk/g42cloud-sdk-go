package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
