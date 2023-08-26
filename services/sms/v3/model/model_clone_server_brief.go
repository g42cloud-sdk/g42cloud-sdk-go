package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CloneServerBrief struct {
	VmId *string `json:"vm_id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o CloneServerBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneServerBrief struct{}"
	}

	return strings.Join([]string{"CloneServerBrief", string(data)}, " ")
}
