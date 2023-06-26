package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PaneSetting struct {
	PaneId string `json:"pane_id"`

	X string `json:"x"`

	Y string `json:"y"`

	Width string `json:"width"`

	Height string `json:"height"`
}

func (o PaneSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PaneSetting struct{}"
	}

	return strings.Join([]string{"PaneSetting", string(data)}, " ")
}
