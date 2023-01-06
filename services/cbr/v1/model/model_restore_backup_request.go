package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RestoreBackupRequest struct {
	BackupId string `json:"backup_id"`

	Body *BackupRestoreReq `json:"body,omitempty"`
}

func (o RestoreBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBackupRequest struct{}"
	}

	return strings.Join([]string{"RestoreBackupRequest", string(data)}, " ")
}
