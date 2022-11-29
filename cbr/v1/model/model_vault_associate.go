package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultAssociate struct {
	DestinationVaultId *string `json:"destination_vault_id,omitempty"`

	PolicyId *string `json:"policy_id,omitempty"`

	AddPolicyIds *[]string `json:"add_policy_ids,omitempty"`
}

func (o VaultAssociate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultAssociate struct{}"
	}

	return strings.Join([]string{"VaultAssociate", string(data)}, " ")
}
