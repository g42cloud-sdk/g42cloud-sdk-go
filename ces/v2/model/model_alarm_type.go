package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmType struct {
}

func (o AlarmType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmType struct{}"
	}

	return strings.Join([]string{"AlarmType", string(data)}, " ")
}
