package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDomainMultiCertificatesRequestBodyContent struct {
	DomainName string `json:"domain_name"`

	HttpsSwitch int32 `json:"https_switch"`

	AccessOriginWay *int32 `json:"access_origin_way,omitempty"`

	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`

	Http2 *int32 `json:"http2,omitempty"`

	CertName *string `json:"cert_name,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	CertificateType *int32 `json:"certificate_type,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequestBodyContent", string(data)}, " ")
}
