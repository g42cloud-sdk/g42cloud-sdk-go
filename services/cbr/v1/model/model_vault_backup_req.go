package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VaultBackupReq struct {
	Checkpoint *VaultBackup `json:"checkpoint"`
}

func (o VaultBackupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBackupReq struct{}"
	}

	return strings.Join([]string{"VaultBackupReq", string(data)}, " ")
}
