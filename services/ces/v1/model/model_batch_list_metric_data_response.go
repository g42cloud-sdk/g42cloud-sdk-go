package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchListMetricDataResponse struct {
	Metrics        *[]BatchMetricData `json:"metrics,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o BatchListMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataResponse struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataResponse", string(data)}, " ")
}
