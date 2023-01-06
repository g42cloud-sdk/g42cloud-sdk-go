package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupResponse", string(data)}, " ")
}
