package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type StatSummary struct {
	Value *float32 `json:"value,omitempty"`

	Date *string `json:"date,omitempty"`
}

func (o StatSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatSummary struct{}"
	}

	return strings.Join([]string{"StatSummary", string(data)}, " ")
}
