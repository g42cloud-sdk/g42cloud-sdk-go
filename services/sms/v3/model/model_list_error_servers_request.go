package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListErrorServersRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Offset int32 `json:"offset"`

	Migproject *string `json:"migproject,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListErrorServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListErrorServersRequest struct{}"
	}

	return strings.Join([]string{"ListErrorServersRequest", string(data)}, " ")
}
