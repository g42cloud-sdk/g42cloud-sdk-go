package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BackupSyncRespBody struct {
	BackupId string `json:"backup_id"`

	OperationLogId string `json:"operation_log_id"`
}

func (o BackupSyncRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupSyncRespBody struct{}"
	}

	return strings.Join([]string{"BackupSyncRespBody", string(data)}, " ")
}
