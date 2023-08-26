package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EnableDomainResponse struct {
	Domain         *DomainsWithPort `json:"domain,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o EnableDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDomainResponse struct{}"
	}

	return strings.Join([]string{"EnableDomainResponse", string(data)}, " ")
}
