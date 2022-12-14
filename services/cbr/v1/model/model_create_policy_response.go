package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreatePolicyResponse struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyResponse", string(data)}, " ")
}
