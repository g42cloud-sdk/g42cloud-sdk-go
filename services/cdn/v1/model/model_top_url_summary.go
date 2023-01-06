package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TopUrlSummary struct {
	Url *string `json:"url,omitempty"`

	Value *int64 `json:"value,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	EndTime *int64 `json:"end_time,omitempty"`

	StatType *string `json:"stat_type,omitempty"`
}

func (o TopUrlSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrlSummary struct{}"
	}

	return strings.Join([]string{"TopUrlSummary", string(data)}, " ")
}
