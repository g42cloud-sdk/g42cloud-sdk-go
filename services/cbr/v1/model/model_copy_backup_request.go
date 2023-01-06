package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CopyBackupRequest struct {
	BackupId string `json:"backup_id"`

	Body *BackupReplicateReq `json:"body,omitempty"`
}

func (o CopyBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyBackupRequest struct{}"
	}

	return strings.Join([]string{"CopyBackupRequest", string(data)}, " ")
}
