package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityGroupResponse struct {
	SecurityGroup  *SecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupResponse", string(data)}, " ")
}
