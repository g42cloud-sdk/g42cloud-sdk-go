package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeletePolicyRequest struct {
	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequest", string(data)}, " ")
}
