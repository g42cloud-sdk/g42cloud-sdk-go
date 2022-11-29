package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointReplicateRespbackups struct {
	BackupId string `json:"backup_id"`

	ReplicationRecordId string `json:"replication_record_id"`
}

func (o CheckpointReplicateRespbackups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointReplicateRespbackups struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateRespbackups", string(data)}, " ")
}
