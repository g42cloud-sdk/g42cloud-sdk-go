package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowResponseHeaderRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowResponseHeaderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResponseHeaderRequest struct{}"
	}

	return strings.Join([]string{"ShowResponseHeaderRequest", string(data)}, " ")
}
