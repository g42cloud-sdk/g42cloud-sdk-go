package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBlackWhiteListRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *BlackWhiteListBody `json:"body,omitempty"`
}

func (o UpdateBlackWhiteListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListRequest struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListRequest", string(data)}, " ")
}
