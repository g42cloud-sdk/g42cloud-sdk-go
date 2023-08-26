package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TargetServerById struct {
	VmId *string `json:"vm_id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o TargetServerById) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetServerById struct{}"
	}

	return strings.Join([]string{"TargetServerById", string(data)}, " ")
}
