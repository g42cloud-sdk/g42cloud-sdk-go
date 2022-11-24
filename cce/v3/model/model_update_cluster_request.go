package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClusterRequest struct {
	ClusterId string `json:"cluster_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	Body *ClusterInformation `json:"body,omitempty"`
}

func (o UpdateClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterRequest", string(data)}, " ")
}
