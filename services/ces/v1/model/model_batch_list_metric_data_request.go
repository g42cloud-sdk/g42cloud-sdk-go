package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchListMetricDataRequest struct {
	ContentType string `json:"Content-Type"`

	Body *BatchListMetricDataRequestBody `json:"body,omitempty"`
}

func (o BatchListMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequest struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequest", string(data)}, " ")
}
