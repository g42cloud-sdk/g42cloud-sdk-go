package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScaleProxyRequestBody struct {
	FlavorRef string `json:"flavor_ref"`

	Delay bool `json:"delay"`
}

func (o ScaleProxyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleProxyRequestBody struct{}"
	}

	return strings.Join([]string{"ScaleProxyRequestBody", string(data)}, " ")
}
