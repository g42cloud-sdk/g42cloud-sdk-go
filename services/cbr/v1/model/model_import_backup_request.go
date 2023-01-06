package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ImportBackupRequest struct {
	Body *BackupSyncReq `json:"body,omitempty"`
}

func (o ImportBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBackupRequest struct{}"
	}

	return strings.Join([]string{"ImportBackupRequest", string(data)}, " ")
}
