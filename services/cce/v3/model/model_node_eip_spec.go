package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NodeEipSpec struct {
	Iptype *string `json:"iptype,omitempty"`

	Bandwidth *NodeBandwidth `json:"bandwidth,omitempty"`
}

func (o NodeEipSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeEipSpec struct{}"
	}

	return strings.Join([]string{"NodeEipSpec", string(data)}, " ")
}
