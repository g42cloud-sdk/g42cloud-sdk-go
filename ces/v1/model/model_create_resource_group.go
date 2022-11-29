package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResourceGroup struct {
	Namespace string `json:"namespace"`

	Dimensions []MetricsDimension `json:"dimensions"`
}

func (o CreateResourceGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroup struct{}"
	}

	return strings.Join([]string{"CreateResourceGroup", string(data)}, " ")
}
