package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowBlackWhiteListRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowBlackWhiteListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlackWhiteListRequest struct{}"
	}

	return strings.Join([]string{"ShowBlackWhiteListRequest", string(data)}, " ")
}
