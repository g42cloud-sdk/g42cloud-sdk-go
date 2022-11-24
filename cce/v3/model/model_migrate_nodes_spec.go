package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateNodesSpec struct {
	Os string `json:"os"`

	ExtendParam *MigrateNodeExtendParam `json:"extendParam,omitempty"`

	Login *Login `json:"login"`

	Nodes []NodeItem `json:"nodes"`
}

func (o MigrateNodesSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodesSpec struct{}"
	}

	return strings.Join([]string{"MigrateNodesSpec", string(data)}, " ")
}
