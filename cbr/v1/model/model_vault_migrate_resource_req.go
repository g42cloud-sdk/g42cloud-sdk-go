package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultMigrateResourceReq struct {
	DestinationVaultId string `json:"destination_vault_id"`

	ResourceIds []string `json:"resource_ids"`
}

func (o VaultMigrateResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultMigrateResourceReq struct{}"
	}

	return strings.Join([]string{"VaultMigrateResourceReq", string(data)}, " ")
}
