package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DatapointForBatchMetric struct {
	Max *float64 `json:"max,omitempty"`

	Min *float64 `json:"min,omitempty"`

	Average *float64 `json:"average,omitempty"`

	Sum *float64 `json:"sum,omitempty"`

	Variance *float64 `json:"variance,omitempty"`

	Timestamp int64 `json:"timestamp"`
}

func (o DatapointForBatchMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatapointForBatchMetric struct{}"
	}

	return strings.Join([]string{"DatapointForBatchMetric", string(data)}, " ")
}
