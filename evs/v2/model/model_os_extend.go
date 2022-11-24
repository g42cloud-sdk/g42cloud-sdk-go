package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OsExtend struct {
	NewSize int32 `json:"new_size"`
}

func (o OsExtend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsExtend struct{}"
	}

	return strings.Join([]string{"OsExtend", string(data)}, " ")
}
