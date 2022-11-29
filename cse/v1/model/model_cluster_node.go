package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterNode struct {
	Id *string `json:"id,omitempty"`

	Az *string `json:"az,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Label *string `json:"label,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o ClusterNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNode struct{}"
	}

	return strings.Join([]string{"ClusterNode", string(data)}, " ")
}
