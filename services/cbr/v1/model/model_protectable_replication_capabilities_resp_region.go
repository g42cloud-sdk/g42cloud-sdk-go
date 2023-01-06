package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
