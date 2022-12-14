package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EnableDomainRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o EnableDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDomainRequest struct{}"
	}

	return strings.Join([]string{"EnableDomainRequest", string(data)}, " ")
}
