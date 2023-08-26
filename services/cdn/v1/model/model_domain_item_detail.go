package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DomainItemDetail struct {
	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	StatType *string `json:"stat_type,omitempty"`

	Domains *[]map[string]interface{} `json:"domains,omitempty"`
}

func (o DomainItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainItemDetail struct{}"
	}

	return strings.Join([]string{"DomainItemDetail", string(data)}, " ")
}
