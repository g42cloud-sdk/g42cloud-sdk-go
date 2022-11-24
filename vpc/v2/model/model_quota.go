package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Quota struct {
	Resources []ResourceResult `json:"resources"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
