package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
