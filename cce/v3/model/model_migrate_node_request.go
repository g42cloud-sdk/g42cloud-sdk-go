package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	TargetClusterId string `json:"target_cluster_id"`

	Body *MigrateNodesTask `json:"body,omitempty"`
}

func (o MigrateNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodeRequest struct{}"
	}

	return strings.Join([]string{"MigrateNodeRequest", string(data)}, " ")
}
