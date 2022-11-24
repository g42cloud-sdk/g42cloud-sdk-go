package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCacheRulesRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowCacheRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCacheRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowCacheRulesRequest", string(data)}, " ")
}
