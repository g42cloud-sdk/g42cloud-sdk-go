package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Rules struct {
	RuleType int32 `json:"rule_type"`

	Content *string `json:"content,omitempty"`

	Ttl int32 `json:"ttl"`

	TtlType int32 `json:"ttl_type"`

	Priority int32 `json:"priority"`
}

func (o Rules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rules struct{}"
	}

	return strings.Join([]string{"Rules", string(data)}, " ")
}
