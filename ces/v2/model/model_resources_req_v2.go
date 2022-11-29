package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesReqV2 struct {
	Resources [][]Dimension `json:"resources"`
}

func (o ResourcesReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesReqV2 struct{}"
	}

	return strings.Join([]string{"ResourcesReqV2", string(data)}, " ")
}
