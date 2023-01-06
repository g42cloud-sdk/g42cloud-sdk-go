package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowBackupRequest struct {
	BackupId string `json:"backup_id"`
}

func (o ShowBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupRequest", string(data)}, " ")
}
