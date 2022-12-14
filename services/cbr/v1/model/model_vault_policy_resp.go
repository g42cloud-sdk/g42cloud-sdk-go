package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultPolicyResp struct {
	DestinationVaultId *string `json:"destination_vault_id,omitempty"`

	PolicyId string `json:"policy_id"`

	VaultId string `json:"vault_id"`
}

func (o VaultPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultPolicyResp struct{}"
	}

	return strings.Join([]string{"VaultPolicyResp", string(data)}, " ")
}
