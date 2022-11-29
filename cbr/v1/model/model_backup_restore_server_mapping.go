package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupRestoreServerMapping struct {
	BackupId string `json:"backup_id"`

	VolumeId string `json:"volume_id"`
}

func (o BackupRestoreServerMapping) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestoreServerMapping struct{}"
	}

	return strings.Join([]string{"BackupRestoreServerMapping", string(data)}, " ")
}
