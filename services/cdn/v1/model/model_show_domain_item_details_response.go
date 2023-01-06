package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainItemDetailsResponse struct {
	DomainItemDetails *DomainItemDetail `json:"domain_item_details,omitempty"`
	HttpStatusCode    int               `json:"-"`
}

func (o ShowDomainItemDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainItemDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainItemDetailsResponse", string(data)}, " ")
}
