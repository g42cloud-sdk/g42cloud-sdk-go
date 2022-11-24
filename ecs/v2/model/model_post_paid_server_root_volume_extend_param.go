package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostPaidServerRootVolumeExtendParam struct {
	ResourceSpecCode *string `json:"resourceSpecCode,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	SnapshotId *string `json:"snapshotId,omitempty"`
}

func (o PostPaidServerRootVolumeExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerRootVolumeExtendParam struct{}"
	}

	return strings.Join([]string{"PostPaidServerRootVolumeExtendParam", string(data)}, " ")
}
