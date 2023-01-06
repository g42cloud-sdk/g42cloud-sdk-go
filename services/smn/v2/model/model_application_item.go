package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ApplicationItem struct {
	Name string `json:"name"`

	Platform string `json:"platform"`

	CreateTime string `json:"create_time"`

	ApplicationUrn string `json:"application_urn"`

	ApplicationId string `json:"application_id"`

	Enabled string `json:"enabled"`

	AppleCertificateExpirationDate *string `json:"apple_certificate_expiration_date,omitempty"`
}

func (o ApplicationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationItem struct{}"
	}

	return strings.Join([]string{"ApplicationItem", string(data)}, " ")
}
