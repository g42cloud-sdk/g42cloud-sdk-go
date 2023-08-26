package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Unit struct {
}

func (o Unit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Unit struct{}"
	}

	return strings.Join([]string{"Unit", string(data)}, " ")
}
