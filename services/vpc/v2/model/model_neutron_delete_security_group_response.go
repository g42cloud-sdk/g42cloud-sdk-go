package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NeutronDeleteSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupResponse", string(data)}, " ")
}
