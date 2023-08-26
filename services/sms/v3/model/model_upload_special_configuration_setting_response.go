package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UploadSpecialConfigurationSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UploadSpecialConfigurationSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSpecialConfigurationSettingResponse struct{}"
	}

	return strings.Join([]string{"UploadSpecialConfigurationSettingResponse", string(data)}, " ")
}
