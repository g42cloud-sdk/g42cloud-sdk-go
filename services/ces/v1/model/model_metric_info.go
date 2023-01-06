package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MetricInfo struct {
	Namespace string `json:"namespace"`

	MetricName string `json:"metric_name"`

	Dimensions []MetricsDimension `json:"dimensions"`
}

func (o MetricInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricInfo struct{}"
	}

	return strings.Join([]string{"MetricInfo", string(data)}, " ")
}
