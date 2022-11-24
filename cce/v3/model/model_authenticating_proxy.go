package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthenticatingProxy struct {
	Ca *string `json:"ca,omitempty"`

	Cert *string `json:"cert,omitempty"`

	PrivateKey *string `json:"privateKey,omitempty"`
}

func (o AuthenticatingProxy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthenticatingProxy struct{}"
	}

	return strings.Join([]string{"AuthenticatingProxy", string(data)}, " ")
}
