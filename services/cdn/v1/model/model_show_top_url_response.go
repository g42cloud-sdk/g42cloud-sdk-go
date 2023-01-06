package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowTopUrlResponse struct {
	ServiceArea *string `json:"service_area,omitempty"`

	TopUrlSummary  *[]TopUrlSummary `json:"top_url_summary,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowTopUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopUrlResponse struct{}"
	}

	return strings.Join([]string{"ShowTopUrlResponse", string(data)}, " ")
}
