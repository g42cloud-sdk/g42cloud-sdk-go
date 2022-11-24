package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetBackupPolicyRequestBody struct {
	BackupPolicy *BackupPolicy `json:"backup_policy"`

	ReserveBackups *bool `json:"reserve_backups,omitempty"`
}

func (o SetBackupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetBackupPolicyRequestBody", string(data)}, " ")
}
