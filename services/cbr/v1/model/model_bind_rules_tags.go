package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BindRulesTags struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o BindRulesTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindRulesTags struct{}"
	}

	return strings.Join([]string{"BindRulesTags", string(data)}, " ")
}
