package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteCloudPersistentVolumeClaimsResponse struct {
	ApiVersion *string `json:"apiVersion,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Metadata *PersistentVolumeClaimMetadata `json:"metadata,omitempty"`

	Spec *PersistentVolumeClaimSpec `json:"spec,omitempty"`

	Status         *PersistentVolumeClaimStatus `json:"status,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o DeleteCloudPersistentVolumeClaimsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudPersistentVolumeClaimsResponse struct{}"
	}

	return strings.Join([]string{"DeleteCloudPersistentVolumeClaimsResponse", string(data)}, " ")
}
