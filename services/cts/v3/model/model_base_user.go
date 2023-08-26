package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BaseUser struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o BaseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseUser struct{}"
	}

	return strings.Join([]string{"BaseUser", string(data)}, " ")
}
