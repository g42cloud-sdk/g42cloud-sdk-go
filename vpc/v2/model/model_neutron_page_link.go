package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NeutronPageLink struct {
	Href string `json:"href"`

	Rel string `json:"rel"`
}

func (o NeutronPageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronPageLink struct{}"
	}

	return strings.Join([]string{"NeutronPageLink", string(data)}, " ")
}
