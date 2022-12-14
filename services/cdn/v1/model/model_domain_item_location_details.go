package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DomainItemLocationDetails struct {
	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	StatType *string `json:"stat_type,omitempty"`

	Domains *[]DomainRegion `json:"domains,omitempty"`
}

func (o DomainItemLocationDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainItemLocationDetails struct{}"
	}

	return strings.Join([]string{"DomainItemLocationDetails", string(data)}, " ")
}
