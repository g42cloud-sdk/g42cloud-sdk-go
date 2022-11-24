package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrePaidServerRootVolumeExtendParam struct {
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	SnapshotId *string `json:"snapshotId,omitempty"`
}

func (o PrePaidServerRootVolumeExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerRootVolumeExtendParam struct{}"
	}

	return strings.Join([]string{"PrePaidServerRootVolumeExtendParam", string(data)}, " ")
}
