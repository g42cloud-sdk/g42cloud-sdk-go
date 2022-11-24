package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateOriginHostRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *OriginHostRequest `json:"body,omitempty"`
}

func (o UpdateOriginHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOriginHostRequest struct{}"
	}

	return strings.Join([]string{"UpdateOriginHostRequest", string(data)}, " ")
}
