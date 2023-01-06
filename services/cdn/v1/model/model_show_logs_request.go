package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowLogsRequest struct {
	DomainName string `json:"domain_name"`

	QueryDate int64 `json:"query_date"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNumber *int32 `json:"page_number,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsRequest struct{}"
	}

	return strings.Join([]string{"ShowLogsRequest", string(data)}, " ")
}
