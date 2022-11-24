package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowServerRemoteConsoleRequestBody struct {
	RemoteConsole *GetServerRemoteConsoleOption `json:"remote_console"`
}

func (o ShowServerRemoteConsoleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerRemoteConsoleRequestBody struct{}"
	}

	return strings.Join([]string{"ShowServerRemoteConsoleRequestBody", string(data)}, " ")
}
