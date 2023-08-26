package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowMetricDataResponse struct {
	Datapoints *[]Datapoint `json:"datapoints,omitempty"`

	MetricName     *string `json:"metric_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricDataResponse", string(data)}, " ")
}
