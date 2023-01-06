package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CdnIps struct {
	Ip *string `json:"ip,omitempty"`

	Belongs *bool `json:"belongs,omitempty"`

	Region *string `json:"region,omitempty"`

	Isp *string `json:"isp,omitempty"`

	Platform *string `json:"platform,omitempty"`
}

func (o CdnIps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdnIps struct{}"
	}

	return strings.Join([]string{"CdnIps", string(data)}, " ")
}
