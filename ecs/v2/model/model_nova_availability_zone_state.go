package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaAvailabilityZoneState struct {
	Available bool `json:"available"`
}

func (o NovaAvailabilityZoneState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaAvailabilityZoneState struct{}"
	}

	return strings.Join([]string{"NovaAvailabilityZoneState", string(data)}, " ")
}
