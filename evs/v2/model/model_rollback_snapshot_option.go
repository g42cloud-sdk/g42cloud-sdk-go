package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RollbackSnapshotOption struct {
	Name *string `json:"name,omitempty"`

	VolumeId string `json:"volume_id"`
}

func (o RollbackSnapshotOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotOption struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotOption", string(data)}, " ")
}
