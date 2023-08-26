package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ConfigsGetBody struct {
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`

	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	UrlAuth *UrlAuthGetBody `json:"url_auth,omitempty"`

	Https *HttpGetBody `json:"https,omitempty"`

	Sources *[]SourcesConfig `json:"sources,omitempty"`

	OriginProtocol *string `json:"origin_protocol,omitempty"`

	ForceRedirect *ForceRedirectConfig `json:"force_redirect,omitempty"`

	Compress *Compress `json:"compress,omitempty"`
}

func (o ConfigsGetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigsGetBody struct{}"
	}

	return strings.Join([]string{"ConfigsGetBody", string(data)}, " ")
}
