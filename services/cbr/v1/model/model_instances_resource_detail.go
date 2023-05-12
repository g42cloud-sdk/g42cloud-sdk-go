package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type InstancesResourceDetail struct {
	Vault *Vault `json:"vault,omitempty"`
}

func (o InstancesResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesResourceDetail struct{}"
	}

	return strings.Join([]string{"InstancesResourceDetail", string(data)}, " ")
}
