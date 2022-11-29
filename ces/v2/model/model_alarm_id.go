package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmId struct {
}

func (o AlarmId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmId struct{}"
	}

	return strings.Join([]string{"AlarmId", string(data)}, " ")
}
