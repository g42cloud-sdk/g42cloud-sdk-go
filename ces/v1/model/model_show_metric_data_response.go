package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
