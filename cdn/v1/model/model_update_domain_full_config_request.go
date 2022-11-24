package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainFullConfigRequest struct {
	DomainName string `json:"domain_name"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ModifyDomainConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainFullConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainFullConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainFullConfigRequest", string(data)}, " ")
}
