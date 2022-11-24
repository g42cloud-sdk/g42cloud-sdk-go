package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeServerOsWithoutCloudInitRequestBody struct {
	OsChange *ChangeServerOsWithoutCloudInitOption `json:"os-change"`
}

func (o ChangeServerOsWithoutCloudInitRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithoutCloudInitRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithoutCloudInitRequestBody", string(data)}, " ")
}
