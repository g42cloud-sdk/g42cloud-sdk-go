package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Configs struct {
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`

	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	UrlAuth *UrlAuth `json:"url_auth,omitempty"`

	Https *HttpPutBody `json:"https,omitempty"`

	Sources *[]SourcesConfig `json:"sources,omitempty"`

	OriginProtocol *string `json:"origin_protocol,omitempty"`

	ForceRedirect *ForceRedirectConfig `json:"force_redirect,omitempty"`

	Compress *Compress `json:"compress,omitempty"`
}

func (o Configs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configs struct{}"
	}

	return strings.Join([]string{"Configs", string(data)}, " ")
}
