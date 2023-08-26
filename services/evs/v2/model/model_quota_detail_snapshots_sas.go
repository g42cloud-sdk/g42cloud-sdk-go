package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type QuotaDetailSnapshotsSas struct {
	InUse int32 `json:"in_use"`

	Limit int32 `json:"limit"`

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailSnapshotsSas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailSnapshotsSas struct{}"
	}

	return strings.Join([]string{"QuotaDetailSnapshotsSas", string(data)}, " ")
}
