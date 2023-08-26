package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Quota struct {
	Type *string `json:"type,omitempty"`

	Used *int64 `json:"used,omitempty"`

	Quota *int64 `json:"quota,omitempty"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
