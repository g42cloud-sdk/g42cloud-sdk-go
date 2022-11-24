package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDomainDetailRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDomainDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailRequest", string(data)}, " ")
}
