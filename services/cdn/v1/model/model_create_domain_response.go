package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateDomainResponse struct {
	Domain         *CreateDomainResponseBodyContent `json:"domain,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CreateDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainResponse", string(data)}, " ")
}
