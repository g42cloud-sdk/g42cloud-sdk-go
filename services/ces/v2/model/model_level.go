package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Level struct {
}

func (o Level) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Level struct{}"
	}

	return strings.Join([]string{"Level", string(data)}, " ")
}
