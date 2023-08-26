package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Metric struct {
	Namespace *string `json:"namespace,omitempty"`

	MetricName *string `json:"metric_name,omitempty"`

	Dimensions *[]Dimension `json:"dimensions,omitempty"`
}

func (o Metric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metric struct{}"
	}

	return strings.Join([]string{"Metric", string(data)}, " ")
}
