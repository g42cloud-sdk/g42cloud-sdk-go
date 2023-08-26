package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowReferRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DomainId string `json:"domain_id"`
}

func (o ShowReferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReferRequest struct{}"
	}

	return strings.Join([]string{"ShowReferRequest", string(data)}, " ")
}
