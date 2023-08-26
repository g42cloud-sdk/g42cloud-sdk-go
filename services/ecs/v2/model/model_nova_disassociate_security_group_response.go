package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaDisassociateSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDisassociateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupResponse", string(data)}, " ")
}
