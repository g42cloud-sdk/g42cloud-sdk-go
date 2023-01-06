package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MetaDataForAlarmHistory struct {
	Total int32 `json:"total"`
}

func (o MetaDataForAlarmHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaDataForAlarmHistory struct{}"
	}

	return strings.Join([]string{"MetaDataForAlarmHistory", string(data)}, " ")
}
