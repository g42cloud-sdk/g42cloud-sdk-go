package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSnapshotRequest struct {
	Body *CreateSnapshotRequestBody `json:"body,omitempty"`
}

func (o CreateSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotRequest struct{}"
	}

	return strings.Join([]string{"CreateSnapshotRequest", string(data)}, " ")
}
