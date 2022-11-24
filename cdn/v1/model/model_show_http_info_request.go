package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHttpInfoRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowHttpInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpInfoRequest", string(data)}, " ")
}
