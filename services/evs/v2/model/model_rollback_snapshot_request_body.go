package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RollbackSnapshotRequestBody struct {
	Rollback *RollbackSnapshotOption `json:"rollback"`
}

func (o RollbackSnapshotRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotRequestBody", string(data)}, " ")
}
