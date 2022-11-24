package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFollow302SwitchRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *Follow302StatusRequest `json:"body,omitempty"`
}

func (o UpdateFollow302SwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFollow302SwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateFollow302SwitchRequest", string(data)}, " ")
}
