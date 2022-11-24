package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisableDomainRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DisableDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDomainRequest struct{}"
	}

	return strings.Join([]string{"DisableDomainRequest", string(data)}, " ")
}
