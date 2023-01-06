package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Count struct {
}

func (o Count) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Count struct{}"
	}

	return strings.Join([]string{"Count", string(data)}, " ")
}
