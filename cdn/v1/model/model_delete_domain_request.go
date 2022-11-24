package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDomainRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DeleteDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainRequest", string(data)}, " ")
}
