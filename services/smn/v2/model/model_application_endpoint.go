package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ApplicationEndpoint struct {
	CreateTime string `json:"create_time"`

	EndpointUrn string `json:"endpoint_urn"`

	UserData string `json:"user_data"`

	Enabled string `json:"enabled"`

	Token string `json:"token"`
}

func (o ApplicationEndpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationEndpoint struct{}"
	}

	return strings.Join([]string{"ApplicationEndpoint", string(data)}, " ")
}
