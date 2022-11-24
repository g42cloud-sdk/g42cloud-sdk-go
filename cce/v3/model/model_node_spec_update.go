package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeSpecUpdate struct {
	Taints []Taint `json:"taints"`

	K8sTags map[string]string `json:"k8sTags"`

	UserTags []UserTag `json:"userTags"`
}

func (o NodeSpecUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpecUpdate struct{}"
	}

	return strings.Join([]string{"NodeSpecUpdate", string(data)}, " ")
}
