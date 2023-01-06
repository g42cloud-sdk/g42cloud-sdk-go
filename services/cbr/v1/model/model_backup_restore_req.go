package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BackupRestoreReq struct {
	Restore *BackupRestore `json:"restore"`
}

func (o BackupRestoreReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestoreReq struct{}"
	}

	return strings.Join([]string{"BackupRestoreReq", string(data)}, " ")
}
