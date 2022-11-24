package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerNetworkUpdate struct {
	Cidrs *[]ContainerCidr `json:"cidrs,omitempty"`
}

func (o ContainerNetworkUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNetworkUpdate struct{}"
	}

	return strings.Join([]string{"ContainerNetworkUpdate", string(data)}, " ")
}
