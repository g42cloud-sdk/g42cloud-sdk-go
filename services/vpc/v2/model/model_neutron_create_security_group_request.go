package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronCreateSecurityGroupRequest struct {
	Body *NeutronCreateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRequest", string(data)}, " ")
}
