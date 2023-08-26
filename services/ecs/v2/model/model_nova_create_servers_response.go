package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type NovaCreateServersResponse struct {
	Server         *NovaCreateServersResult `json:"server,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o NovaCreateServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersResponse struct{}"
	}

	return strings.Join([]string{"NovaCreateServersResponse", string(data)}, " ")
}
