package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Computes struct {
	GroupType *string `json:"group_type,omitempty"`

	ComputeFlavors *[]ScaleFlavors `json:"compute_flavors,omitempty"`
}

func (o Computes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Computes struct{}"
	}

	return strings.Join([]string{"Computes", string(data)}, " ")
}
