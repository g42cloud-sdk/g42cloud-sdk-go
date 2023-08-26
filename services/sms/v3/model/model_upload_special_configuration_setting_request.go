package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UploadSpecialConfigurationSettingRequest struct {
	TaskId string `json:"task_id"`

	Body *ConfigurationRequestBody `json:"body,omitempty"`
}

func (o UploadSpecialConfigurationSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSpecialConfigurationSettingRequest struct{}"
	}

	return strings.Join([]string{"UploadSpecialConfigurationSettingRequest", string(data)}, " ")
}
