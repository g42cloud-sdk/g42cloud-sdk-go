package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DisassociateVaultPolicyRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultDissociate `json:"body,omitempty"`
}

func (o DisassociateVaultPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVaultPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisassociateVaultPolicyRequest", string(data)}, " ")
}
