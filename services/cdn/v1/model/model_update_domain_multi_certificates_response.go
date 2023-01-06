package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDomainMultiCertificatesResponse struct {
	Https          *UpdateDomainMultiCertificatesResponseBodyContent `json:"https,omitempty"`
	HttpStatusCode int                                               `json:"-"`
}

func (o UpdateDomainMultiCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesResponse", string(data)}, " ")
}
