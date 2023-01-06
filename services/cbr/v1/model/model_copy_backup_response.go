package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CopyBackupResponse struct {
	Replication    *BackupReplicateRespBody `json:"replication,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CopyBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyBackupResponse struct{}"
	}

	return strings.Join([]string{"CopyBackupResponse", string(data)}, " ")
}
