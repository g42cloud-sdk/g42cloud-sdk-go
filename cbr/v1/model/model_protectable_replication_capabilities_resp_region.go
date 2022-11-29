package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectableReplicationCapabilitiesRespRegion struct {
	Name string `json:"name"`

	ReplicationDestinations []string `json:"replication_destinations"`
}

func (o ProtectableReplicationCapabilitiesRespRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableReplicationCapabilitiesRespRegion struct{}"
	}

	return strings.Join([]string{"ProtectableReplicationCapabilitiesRespRegion", string(data)}, " ")
}
