package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceRequest struct {
	Kind string `json:"kind"`

	ApiVersion string `json:"apiVersion"`

	Metadata *Metadata `json:"metadata"`

	Spec *InstanceRequestSpec `json:"spec"`
}

func (o InstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRequest struct{}"
	}

	return strings.Join([]string{"InstanceRequest", string(data)}, " ")
}
