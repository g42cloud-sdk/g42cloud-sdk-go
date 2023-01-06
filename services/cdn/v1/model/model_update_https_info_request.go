package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateHttpsInfoRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *HttpInfoRequest `json:"body,omitempty"`
}

func (o UpdateHttpsInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpsInfoRequest", string(data)}, " ")
}
