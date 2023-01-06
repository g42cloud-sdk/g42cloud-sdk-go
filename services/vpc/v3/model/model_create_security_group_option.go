package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupOption struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateSecurityGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupOption", string(data)}, " ")
}
