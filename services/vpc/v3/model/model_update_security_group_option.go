package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateSecurityGroupOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateSecurityGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupOption", string(data)}, " ")
}
