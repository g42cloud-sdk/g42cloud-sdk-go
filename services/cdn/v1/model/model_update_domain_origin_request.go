package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDomainOriginRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *OriginRequest `json:"body,omitempty"`
}

func (o UpdateDomainOriginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainOriginRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainOriginRequest", string(data)}, " ")
}
