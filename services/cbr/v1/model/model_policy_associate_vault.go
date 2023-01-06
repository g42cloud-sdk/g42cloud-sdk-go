package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PolicyAssociateVault struct {
	DestinationVaultId *string `json:"destination_vault_id,omitempty"`

	VaultId string `json:"vault_id"`
}

func (o PolicyAssociateVault) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAssociateVault struct{}"
	}

	return strings.Join([]string{"PolicyAssociateVault", string(data)}, " ")
}
