package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpPutBody struct {
	HttpsStatus *string `json:"https_status,omitempty"`

	CertificateName *string `json:"certificate_name,omitempty"`

	CertificateValue *string `json:"certificate_value,omitempty"`

	PrivateKey *string `json:"private_key,omitempty"`

	CertificateSource *int32 `json:"certificate_source,omitempty"`

	Http2Status *string `json:"http2_status,omitempty"`
}

func (o HttpPutBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpPutBody struct{}"
	}

	return strings.Join([]string{"HttpPutBody", string(data)}, " ")
}
