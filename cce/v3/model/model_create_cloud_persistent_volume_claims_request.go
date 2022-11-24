package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCloudPersistentVolumeClaimsRequest struct {
	Namespace string `json:"namespace"`

	XClusterID *string `json:"X-Cluster-ID,omitempty"`

	Body *PersistentVolumeClaim `json:"body,omitempty"`
}

func (o CreateCloudPersistentVolumeClaimsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPersistentVolumeClaimsRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudPersistentVolumeClaimsRequest", string(data)}, " ")
}
