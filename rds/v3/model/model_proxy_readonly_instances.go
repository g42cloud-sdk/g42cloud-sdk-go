package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyReadonlyInstances struct {
	Id string `json:"id"`

	Weight int32 `json:"weight"`
}

func (o ProxyReadonlyInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyReadonlyInstances struct{}"
	}

	return strings.Join([]string{"ProxyReadonlyInstances", string(data)}, " ")
}
