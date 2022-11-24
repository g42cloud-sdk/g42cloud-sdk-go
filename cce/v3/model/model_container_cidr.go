package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerCidr struct {
	Cidr string `json:"cidr"`
}

func (o ContainerCidr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerCidr struct{}"
	}

	return strings.Join([]string{"ContainerCidr", string(data)}, " ")
}
