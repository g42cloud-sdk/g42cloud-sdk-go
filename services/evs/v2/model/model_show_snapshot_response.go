package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSnapshotResponse struct {
	Snapshot       *SnapshotDetails `json:"snapshot,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowSnapshotResponse", string(data)}, " ")
}
