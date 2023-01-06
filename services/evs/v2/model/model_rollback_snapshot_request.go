package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RollbackSnapshotRequest struct {
	SnapshotId string `json:"snapshot_id"`

	Body *RollbackSnapshotRequestBody `json:"body,omitempty"`
}

func (o RollbackSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotRequest struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotRequest", string(data)}, " ")
}
