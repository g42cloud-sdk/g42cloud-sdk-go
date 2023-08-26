package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListSnapshotsResponse struct {
	Count *int32 `json:"count,omitempty"`

	Snapshots *[]SnapshotList `json:"snapshots,omitempty"`

	SnapshotsLinks *[]Link `json:"snapshots_links,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotsResponse", string(data)}, " ")
}
