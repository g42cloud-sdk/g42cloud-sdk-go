package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateDomainRequest struct {
	Body *CreateDomainRequestBody `json:"body,omitempty"`
}

func (o CreateDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainRequest", string(data)}, " ")
}
