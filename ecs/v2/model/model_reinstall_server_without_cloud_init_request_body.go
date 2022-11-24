package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReinstallServerWithoutCloudInitRequestBody struct {
	OsReinstall *ReinstallServerWithoutCloudInitOption `json:"os-reinstall"`
}

func (o ReinstallServerWithoutCloudInitRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithoutCloudInitRequestBody struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithoutCloudInitRequestBody", string(data)}, " ")
}
