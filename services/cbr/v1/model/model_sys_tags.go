package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SysTags struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o SysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTags struct{}"
	}

	return strings.Join([]string{"SysTags", string(data)}, " ")
}
