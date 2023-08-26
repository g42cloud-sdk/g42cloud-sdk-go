package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AssociateVaultPolicyResponse struct {
	AssociatePolicy *VaultPolicyResp `json:"associate_policy,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o AssociateVaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateVaultPolicyResponse", string(data)}, " ")
}
