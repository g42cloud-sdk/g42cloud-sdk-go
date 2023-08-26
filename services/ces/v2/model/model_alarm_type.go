package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
