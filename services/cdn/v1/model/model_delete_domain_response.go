package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteDomainResponse struct {
	Domain         *DomainsWithPort `json:"domain,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DeleteDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainResponse", string(data)}, " ")
}
