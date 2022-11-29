package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricName struct {
}

func (o MetricName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricName struct{}"
	}

	return strings.Join([]string{"MetricName", string(data)}, " ")
}
