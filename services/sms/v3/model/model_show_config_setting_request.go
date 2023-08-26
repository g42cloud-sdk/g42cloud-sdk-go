package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowConfigSettingRequest struct {
	TaskId string `json:"task_id"`

	ConfigKey *string `json:"config_key,omitempty"`
}

func (o ShowConfigSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigSettingRequest", string(data)}, " ")
}
