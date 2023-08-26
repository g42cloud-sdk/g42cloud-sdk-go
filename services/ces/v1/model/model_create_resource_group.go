package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
