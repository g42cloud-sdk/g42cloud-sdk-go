package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainMultiCertificatesResponseBodyContent struct {
	DomainName string `json:"domain_name"`

	HttpsSwitch *int32 `json:"https_switch,omitempty"`

	AccessOriginWay *int32 `json:"access_origin_way,omitempty"`

	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`

	Http2 *int32 `json:"http2,omitempty"`

	CertName *string `json:"cert_name,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	CertificateType *int32 `json:"certificate_type,omitempty"`

	ExpirationTime *int64 `json:"expiration_time,omitempty"`
}

func (o UpdateDomainMultiCertificatesResponseBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesResponseBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesResponseBodyContent", string(data)}, " ")
}
