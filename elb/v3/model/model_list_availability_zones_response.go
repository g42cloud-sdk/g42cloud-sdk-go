package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAvailabilityZonesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	AvailabilityZones *[][]AvailabilityZone `json:"availability_zones,omitempty"`
	HttpStatusCode    int                   `json:"-"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
