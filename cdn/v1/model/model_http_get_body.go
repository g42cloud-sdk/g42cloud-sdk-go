package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpGetBody struct {
	HttpsStatus *string `json:"https_status,omitempty"`

	CertificateName *string `json:"certificate_name,omitempty"`

	CertificateValue *string `json:"certificate_value,omitempty"`

	CertificateSource *int32 `json:"certificate_source,omitempty"`

	Http2Status *string `json:"http2_status,omitempty"`
}

func (o HttpGetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpGetBody struct{}"
	}

	return strings.Join([]string{"HttpGetBody", string(data)}, " ")
}
