package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSnapshotRequestBody struct {
	Snapshot *CreateSnapshotOption `json:"snapshot"`
}

func (o CreateSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSnapshotRequestBody", string(data)}, " ")
}
