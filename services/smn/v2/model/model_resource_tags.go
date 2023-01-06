package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResourceTags struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o ResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTags struct{}"
	}

	return strings.Join([]string{"ResourceTags", string(data)}, " ")
}
