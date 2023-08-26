package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCertificatesHttpsInfoResponse struct {
	Total *int32 `json:"total,omitempty"`

	Https          *[]HttpsDetail `json:"https,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowCertificatesHttpsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificatesHttpsInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificatesHttpsInfoResponse", string(data)}, " ")
}
