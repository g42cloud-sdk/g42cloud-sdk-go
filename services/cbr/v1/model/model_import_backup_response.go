package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ImportBackupResponse struct {
	Sync           *[]BackupSyncRespBody `json:"sync,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ImportBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBackupResponse struct{}"
	}

	return strings.Join([]string{"ImportBackupResponse", string(data)}, " ")
}
