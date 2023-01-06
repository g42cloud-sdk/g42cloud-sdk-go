package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MigrateVaultResourceRequest struct {
	VaultId string `json:"vault_id"`

	Body *VaultMigrateResourceReq `json:"body,omitempty"`
}

func (o MigrateVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"MigrateVaultResourceRequest", string(data)}, " ")
}
