package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type GetRestoreTimeResponseRestoreTime struct {
	StartTime int32 `json:"start_time"`

	EndTime int32 `json:"end_time"`
}

func (o GetRestoreTimeResponseRestoreTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetRestoreTimeResponseRestoreTime struct{}"
	}

	return strings.Join([]string{"GetRestoreTimeResponseRestoreTime", string(data)}, " ")
}
