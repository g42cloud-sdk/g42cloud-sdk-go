package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigSettingResponse struct {
	TaskId *string `json:"task_id,omitempty"`

	MigrateType *string `json:"migrate_type,omitempty"`

	Configurations *string `json:"configurations,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConfigSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigSettingResponse", string(data)}, " ")
}
