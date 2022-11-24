package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowInstanceConfigurationRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceConfigurationRequest", string(data)}, " ")
}
