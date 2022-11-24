package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
