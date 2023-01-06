package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AlarmEnabled struct {
}

func (o AlarmEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmEnabled struct{}"
	}

	return strings.Join([]string{"AlarmEnabled", string(data)}, " ")
}
