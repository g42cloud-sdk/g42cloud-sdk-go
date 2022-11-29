package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
