package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApplicationAttributesResponseBodyAttributes struct {
	Enabled string `json:"enabled"`

	AppleCertificateExpirationDate *string `json:"apple_certificate_expiration_date,omitempty"`
}

func (o ListApplicationAttributesResponseBodyAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationAttributesResponseBodyAttributes struct{}"
	}

	return strings.Join([]string{"ListApplicationAttributesResponseBodyAttributes", string(data)}, " ")
}
