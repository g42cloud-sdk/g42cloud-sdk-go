package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BackupSyncReq struct {
	Sync []BackupSync `json:"sync"`
}

func (o BackupSyncReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupSyncReq struct{}"
	}

	return strings.Join([]string{"BackupSyncReq", string(data)}, " ")
}
