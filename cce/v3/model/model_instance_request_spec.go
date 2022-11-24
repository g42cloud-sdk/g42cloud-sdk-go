package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceRequestSpec struct {
	Version string `json:"version"`

	ClusterID string `json:"clusterID"`

	Values map[string]interface{} `json:"values"`

	AddonTemplateName string `json:"addonTemplateName"`
}

func (o InstanceRequestSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceRequestSpec struct{}"
	}

	return strings.Join([]string{"InstanceRequestSpec", string(data)}, " ")
}
