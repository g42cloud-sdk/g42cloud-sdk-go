package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SourcesConfig struct {
	OriginAddr string `json:"origin_addr"`

	OriginType string `json:"origin_type"`

	Priority int32 `json:"priority"`

	ObsWebHostingStatus *string `json:"obs_web_hosting_status,omitempty"`

	HttpPort *int32 `json:"http_port,omitempty"`

	HttpsPort *int32 `json:"https_port,omitempty"`

	HostName *string `json:"host_name,omitempty"`
}

func (o SourcesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourcesConfig struct{}"
	}

	return strings.Join([]string{"SourcesConfig", string(data)}, " ")
}
