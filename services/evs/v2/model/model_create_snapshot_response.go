package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateSnapshotResponse struct {
	Snapshot       *SnapshotDetails `json:"snapshot,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CreateSnapshotResponse", string(data)}, " ")
}
