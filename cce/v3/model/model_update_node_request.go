package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	NodeId string `json:"node_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	Body *ClusterNodeInformation `json:"body,omitempty"`
}

func (o UpdateNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodeRequest", string(data)}, " ")
}
