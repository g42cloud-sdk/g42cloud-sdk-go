package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaAvailabilityZone struct {
	Hosts []string `json:"hosts"`

	ZoneName string `json:"zoneName"`

	ZoneState *NovaAvailabilityZoneState `json:"zoneState"`
}

func (o NovaAvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAvailabilityZone struct{}"
	}

	return strings.Join([]string{"NovaAvailabilityZone", string(data)}, " ")
}
