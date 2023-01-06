package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteBackupRequest struct {
	BackupId string `json:"backup_id"`
}

func (o DeleteBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupRequest", string(data)}, " ")
}
