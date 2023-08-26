package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListStatSummaryRequest struct {
	StartTime string `json:"start_time"`

	EndTime string `json:"end_time"`

	StatType string `json:"stat_type"`
}

func (o ListStatSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatSummaryRequest struct{}"
	}

	return strings.Join([]string{"ListStatSummaryRequest", string(data)}, " ")
}
