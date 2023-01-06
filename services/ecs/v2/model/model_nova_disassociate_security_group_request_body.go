package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaDisassociateSecurityGroupRequestBody struct {
	RemoveSecurityGroup *NovaRemoveSecurityGroupOption `json:"removeSecurityGroup"`
}

func (o NovaDisassociateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupRequestBody", string(data)}, " ")
}