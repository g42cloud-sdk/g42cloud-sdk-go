package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoSync struct {
	SyncBackupNum *int32 `json:"sync_backup_num,omitempty"`

	DeleteBackupNum *int32 `json:"delete_backup_num,omitempty"`

	ErrSyncBackupNum *int32 `json:"err_sync_backup_num,omitempty"`
}

func (o OpExtendInfoSync) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoSync struct{}"
	}

	return strings.Join([]string{"OpExtendInfoSync", string(data)}, " ")
}
