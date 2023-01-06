package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaAssociateSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaAssociateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAssociateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NovaAssociateSecurityGroupResponse", string(data)}, " ")
}
