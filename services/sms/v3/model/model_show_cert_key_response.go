package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCertKeyResponse struct {
	Cert *string `json:"cert,omitempty"`

	PrivateKey     *string `json:"private_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCertKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowCertKeyResponse", string(data)}, " ")
}
