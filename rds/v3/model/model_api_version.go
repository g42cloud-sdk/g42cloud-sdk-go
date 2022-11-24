package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiVersion struct {
	Id string `json:"id"`

	Links []LinksInfoResponse `json:"links"`

	Status string `json:"status"`

	Updated string `json:"updated"`
}

func (o ApiVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersion struct{}"
	}

	return strings.Join([]string{"ApiVersion", string(data)}, " ")
}
