package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSnapshotRequest struct {
	SnapshotId string `json:"snapshot_id"`
}

func (o ShowSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowSnapshotRequest", string(data)}, " ")
}
