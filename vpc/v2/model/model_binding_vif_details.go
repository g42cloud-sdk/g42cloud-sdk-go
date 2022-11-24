package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindingVifDetails struct {
	PrimaryInterface *bool `json:"primary_interface,omitempty"`

	PortFilter *bool `json:"port_filter,omitempty"`

	OvsHybridPlug *bool `json:"ovs_hybrid_plug,omitempty"`
}

func (o BindingVifDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingVifDetails struct{}"
	}

	return strings.Join([]string{"BindingVifDetails", string(data)}, " ")
}
