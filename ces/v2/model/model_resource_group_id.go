package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceGroupId struct {
}

func (o ResourceGroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupId struct{}"
	}

	return strings.Join([]string{"ResourceGroupId", string(data)}, " ")
}
