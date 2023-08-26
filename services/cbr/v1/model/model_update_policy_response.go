package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdatePolicyResponse struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyResponse", string(data)}, " ")
}
