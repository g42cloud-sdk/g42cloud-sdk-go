package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowOriginHostRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowOriginHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOriginHostRequest struct{}"
	}

	return strings.Join([]string{"ShowOriginHostRequest", string(data)}, " ")
}
