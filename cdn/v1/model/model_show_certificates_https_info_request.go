package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCertificatesHttpsInfoRequest struct {
	PageSize *int32 `json:"page_size,omitempty"`

	PageNumber *int32 `json:"page_number,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	UserDomainId *string `json:"user_domain_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowCertificatesHttpsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificatesHttpsInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificatesHttpsInfoRequest", string(data)}, " ")
}
