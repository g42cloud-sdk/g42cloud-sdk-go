package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DisassociateVaultPolicyResponse struct {
	DissociatePolicy *VaultPolicyResp `json:"dissociate_policy,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o DisassociateVaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisassociateVaultPolicyResponse", string(data)}, " ")
}
