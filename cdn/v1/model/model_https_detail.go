package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpsDetail struct {
	DomainId *string `json:"domain_id,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	CertName *string `json:"cert_name,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	CertificateType *int32 `json:"certificate_type,omitempty"`

	ExpirationTime *int64 `json:"expiration_time,omitempty"`

	HttpsStatus *int32 `json:"https_status,omitempty"`

	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`

	Http2 *int32 `json:"http2,omitempty"`
}

func (o HttpsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpsDetail struct{}"
	}

	return strings.Join([]string{"HttpsDetail", string(data)}, " ")
}
