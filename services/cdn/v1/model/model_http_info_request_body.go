package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type HttpInfoRequestBody struct {
	CertName string `json:"cert_name"`

	HttpsStatus int32 `json:"https_status"`

	Certificate *string `json:"certificate,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	Http2 *int32 `json:"http2,omitempty"`

	CertificateType *int32 `json:"certificate_type,omitempty"`

	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`
}

func (o HttpInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpInfoRequestBody struct{}"
	}

	return strings.Join([]string{"HttpInfoRequestBody", string(data)}, " ")
}
