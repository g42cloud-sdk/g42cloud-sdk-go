package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainMultiCertificatesRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *UpdateDomainMultiCertificatesRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequest", string(data)}, " ")
}
