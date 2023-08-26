package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdatePolicyRequest struct {
	PolicyId string `json:"policy_id"`

	Body *PolicyUpdateReq `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}
