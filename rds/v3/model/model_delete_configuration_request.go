package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteConfigurationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	ConfigId string `json:"config_id"`
}

func (o DeleteConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigurationRequest", string(data)}, " ")
}
