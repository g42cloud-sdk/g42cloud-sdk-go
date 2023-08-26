package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
