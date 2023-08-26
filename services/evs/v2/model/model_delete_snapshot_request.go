package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSnapshotRequest struct {
	SnapshotId string `json:"snapshot_id"`
}

func (o DeleteSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotRequest struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotRequest", string(data)}, " ")
}
