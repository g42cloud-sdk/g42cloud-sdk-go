package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TagMatch struct {
	Key string `json:"key"`

	Value string `json:"value"`
}

func (o TagMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMatch struct{}"
	}

	return strings.Join([]string{"TagMatch", string(data)}, " ")
}
