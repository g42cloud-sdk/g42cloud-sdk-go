package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DomainOriginHost struct {
	DomainId *string `json:"domain_id,omitempty"`

	OriginHostType string `json:"origin_host_type"`

	CustomizeDomain *string `json:"customize_domain,omitempty"`
}

func (o DomainOriginHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainOriginHost struct{}"
	}

	return strings.Join([]string{"DomainOriginHost", string(data)}, " ")
}
