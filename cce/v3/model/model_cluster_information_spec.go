package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterInformationSpec struct {
	Description *string `json:"description,omitempty"`

	CustomSan *[]string `json:"customSan,omitempty"`

	ContainerNetwork *ContainerNetworkUpdate `json:"containerNetwork,omitempty"`
}

func (o ClusterInformationSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInformationSpec struct{}"
	}

	return strings.Join([]string{"ClusterInformationSpec", string(data)}, " ")
}
