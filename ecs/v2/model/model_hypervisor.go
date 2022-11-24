package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Hypervisor struct {
	HypervisorType *string `json:"hypervisor_type,omitempty"`

	CsdHypervisor *string `json:"csd_hypervisor,omitempty"`
}

func (o Hypervisor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Hypervisor struct{}"
	}

	return strings.Join([]string{"Hypervisor", string(data)}, " ")
}
