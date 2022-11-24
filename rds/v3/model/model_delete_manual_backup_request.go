package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteManualBackupRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	BackupId string `json:"backup_id"`
}

func (o DeleteManualBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManualBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteManualBackupRequest", string(data)}, " ")
}
