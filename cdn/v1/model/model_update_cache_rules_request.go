package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCacheRulesRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CacheConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateCacheRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCacheRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateCacheRulesRequest", string(data)}, " ")
}
