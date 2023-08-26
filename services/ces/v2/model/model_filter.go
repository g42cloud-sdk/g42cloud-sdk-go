package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Filter struct {
}

func (o Filter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Filter struct{}"
	}

	return strings.Join([]string{"Filter", string(data)}, " ")
}
