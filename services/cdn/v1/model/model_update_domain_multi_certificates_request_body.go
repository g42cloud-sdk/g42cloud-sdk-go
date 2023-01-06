package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDomainMultiCertificatesRequestBody struct {
	Https *UpdateDomainMultiCertificatesRequestBodyContent `json:"https,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequestBody", string(data)}, " ")
}
