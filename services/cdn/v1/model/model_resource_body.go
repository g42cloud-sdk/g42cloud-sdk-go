package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResourceBody struct {
	Sources []SourceWithPort `json:"sources"`
}

func (o ResourceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceBody struct{}"
	}

	return strings.Join([]string{"ResourceBody", string(data)}, " ")
}
