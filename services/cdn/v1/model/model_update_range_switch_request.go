package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateRangeSwitchRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DomainId string `json:"domain_id"`

	Body *RangeStatusRequest `json:"body,omitempty"`
}

func (o UpdateRangeSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRangeSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateRangeSwitchRequest", string(data)}, " ")
}
