package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainItemLocationDetailsResponse struct {
	DomainItemLocationDetails *DomainItemLocationDetails `json:"domain_item_location_details,omitempty"`
	HttpStatusCode            int                        `json:"-"`
}

func (o ShowDomainItemLocationDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainItemLocationDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainItemLocationDetailsResponse", string(data)}, " ")
}
