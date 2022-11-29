package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricsDimension struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

func (o MetricsDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricsDimension struct{}"
	}

	return strings.Join([]string{"MetricsDimension", string(data)}, " ")
}
