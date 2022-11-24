package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateReferRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DomainId string `json:"domain_id"`

	Body *RefererBody `json:"body,omitempty"`
}

func (o UpdateReferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReferRequest struct{}"
	}

	return strings.Join([]string{"UpdateReferRequest", string(data)}, " ")
}
