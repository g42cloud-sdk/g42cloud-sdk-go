package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSecurityPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityPolicyResponse", string(data)}, " ")
}
