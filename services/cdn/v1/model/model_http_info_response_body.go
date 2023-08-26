package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type HttpInfoResponseBody struct {
	HttpsStatus *int32 `json:"https_status,omitempty"`

	CertName *string `json:"cert_name,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	CertificateType *int32 `json:"certificate_type,omitempty"`

	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`

	Http2 *int32 `json:"http2,omitempty"`

	ExpirationTime *int64 `json:"expiration_time,omitempty"`
}

func (o HttpInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpInfoResponseBody struct{}"
	}

	return strings.Join([]string{"HttpInfoResponseBody", string(data)}, " ")
}
