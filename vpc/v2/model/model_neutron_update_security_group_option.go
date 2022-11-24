package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronUpdateSecurityGroupOption struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o NeutronUpdateSecurityGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupOption", string(data)}, " ")
}
