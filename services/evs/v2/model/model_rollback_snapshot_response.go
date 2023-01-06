package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RollbackSnapshotResponse struct {
	Rollback       *RollbackInfo `json:"rollback,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o RollbackSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackSnapshotResponse struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotResponse", string(data)}, " ")
}
