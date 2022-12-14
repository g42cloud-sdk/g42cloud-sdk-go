package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainFullConfigRequest struct {
	DomainName string `json:"domain_name"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainFullConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainFullConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainFullConfigRequest", string(data)}, " ")
}
