package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageLink struct {
	Href string `json:"href"`

	Rel string `json:"rel"`
}

func (o PageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageLink struct{}"
	}

	return strings.Join([]string{"PageLink", string(data)}, " ")
}
