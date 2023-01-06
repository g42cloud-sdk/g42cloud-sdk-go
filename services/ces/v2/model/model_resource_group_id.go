package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
