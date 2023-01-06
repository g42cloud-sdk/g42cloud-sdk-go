package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultBackup struct {
	Parameters *CheckpointParam `json:"parameters,omitempty"`

	VaultId string `json:"vault_id"`
}

func (o VaultBackup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBackup struct{}"
	}

	return strings.Join([]string{"VaultBackup", string(data)}, " ")
}
