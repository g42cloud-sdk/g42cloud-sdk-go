package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddonInstance struct {
	Kind string `json:"kind"`

	ApiVersion string `json:"apiVersion"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *InstanceSpec `json:"spec"`

	Status *AddonInstanceStatus `json:"status"`
}

func (o AddonInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonInstance struct{}"
	}

	return strings.Join([]string{"AddonInstance", string(data)}, " ")
}
