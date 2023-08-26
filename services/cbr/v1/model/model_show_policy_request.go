package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowPolicyRequest struct {
	PolicyId string `json:"policy_id"`
}

func (o ShowPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyRequest", string(data)}, " ")
}
