package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Resource struct {
	Type string `json:"type"`

	Used int32 `json:"used"`

	Unit string `json:"unit"`

	Quota int32 `json:"quota"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
