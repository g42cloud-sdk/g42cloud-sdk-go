package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	SecurityGroup *CreateSecurityGroupOption `json:"security_group"`
}

func (o CreateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRequestBody", string(data)}, " ")
}
