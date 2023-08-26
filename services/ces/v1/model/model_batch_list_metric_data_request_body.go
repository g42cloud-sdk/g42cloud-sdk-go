package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchListMetricDataRequestBody struct {
	Metrics []MetricInfo `json:"metrics"`

	Period string `json:"period"`

	Filter string `json:"filter"`

	From int64 `json:"from"`

	To int64 `json:"to"`
}

func (o BatchListMetricDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequestBody", string(data)}, " ")
}
