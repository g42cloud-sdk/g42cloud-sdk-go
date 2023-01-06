package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateResponseHeaderRequest struct {
	DomainId string `json:"domain_id"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *HeaderBody `json:"body,omitempty"`
}

func (o UpdateResponseHeaderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResponseHeaderRequest struct{}"
	}

	return strings.Join([]string{"UpdateResponseHeaderRequest", string(data)}, " ")
}
