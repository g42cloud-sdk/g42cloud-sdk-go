package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistentVolumeClaim struct {
	ApiVersion string `json:"apiVersion"`

	Kind string `json:"kind"`

	Metadata *PersistentVolumeClaimMetadata `json:"metadata"`

	Spec *PersistentVolumeClaimSpec `json:"spec"`

	Status *PersistentVolumeClaimStatus `json:"status,omitempty"`
}

func (o PersistentVolumeClaim) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaim struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaim", string(data)}, " ")
}
