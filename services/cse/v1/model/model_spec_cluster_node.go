package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SpecClusterNode struct {
	ClusterNodes *[]ClusterNode `json:"cluster_nodes,omitempty"`
}

func (o SpecClusterNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecClusterNode struct{}"
	}

	return strings.Join([]string{"SpecClusterNode", string(data)}, " ")
}
