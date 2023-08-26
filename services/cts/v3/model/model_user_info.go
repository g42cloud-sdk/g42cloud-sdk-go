package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UserInfo struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Domain *BaseUser `json:"domain,omitempty"`
}

func (o UserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserInfo struct{}"
	}

	return strings.Join([]string{"UserInfo", string(data)}, " ")
}
