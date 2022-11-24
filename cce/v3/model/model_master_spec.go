package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MasterSpec struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	FaultDomain *string `json:"faultDomain,omitempty"`
}

func (o MasterSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterSpec struct{}"
	}

	return strings.Join([]string{"MasterSpec", string(data)}, " ")
}
