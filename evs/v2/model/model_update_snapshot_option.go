package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotOption struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o UpdateSnapshotOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotOption struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotOption", string(data)}, " ")
}
