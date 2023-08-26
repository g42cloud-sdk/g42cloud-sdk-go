package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowDomainFullConfigResponse struct {
	Configs        *ConfigsGetBody `json:"configs,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowDomainFullConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainFullConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainFullConfigResponse", string(data)}, " ")
}
