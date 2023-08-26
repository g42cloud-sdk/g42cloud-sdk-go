package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMetricDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetricDataResponse struct{}"
	}

	return strings.Join([]string{"CreateMetricDataResponse", string(data)}, " ")
}
