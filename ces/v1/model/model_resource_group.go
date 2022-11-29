package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceGroup struct {
	Namespace *string `json:"namespace,omitempty"`

	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o ResourceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroup struct{}"
	}

	return strings.Join([]string{"ResourceGroup", string(data)}, " ")
}
