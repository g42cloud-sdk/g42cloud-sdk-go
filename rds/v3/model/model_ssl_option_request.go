package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SslOptionRequest struct {
	SslOption bool `json:"ssl_option"`
}

func (o SslOptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SslOptionRequest struct{}"
	}

	return strings.Join([]string{"SslOptionRequest", string(data)}, " ")
}
