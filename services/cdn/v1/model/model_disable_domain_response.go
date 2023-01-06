package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DisableDomainResponse struct {
	Domain         *DomainsWithPort `json:"domain,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DisableDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDomainResponse struct{}"
	}

	return strings.Join([]string{"DisableDomainResponse", string(data)}, " ")
}
