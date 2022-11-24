package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnchangeableParam struct {
	LowerCaseTableNames *string `json:"lower_case_table_names,omitempty"`
}

func (o UnchangeableParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnchangeableParam struct{}"
	}

	return strings.Join([]string{"UnchangeableParam", string(data)}, " ")
}
