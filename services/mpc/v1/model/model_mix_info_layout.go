package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MixInfoLayout struct {
	Panes []PaneSetting `json:"panes"`
}

func (o MixInfoLayout) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MixInfoLayout struct{}"
	}

	return strings.Join([]string{"MixInfoLayout", string(data)}, " ")
}
