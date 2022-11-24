package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistentVolumeClaimMetadata struct {
	Name string `json:"name"`

	Labels *string `json:"labels,omitempty"`
}

func (o PersistentVolumeClaimMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaimMetadata struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaimMetadata", string(data)}, " ")
}
