package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type FlavorBrief struct {
	Flavor *string `json:"flavor,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o FlavorBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorBrief struct{}"
	}

	return strings.Join([]string{"FlavorBrief", string(data)}, " ")
}
