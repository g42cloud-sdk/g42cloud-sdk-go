package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpcOption struct {
	Cidr *string `json:"cidr,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateVpcOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcOption struct{}"
	}

	return strings.Join([]string{"CreateVpcOption", string(data)}, " ")
}
