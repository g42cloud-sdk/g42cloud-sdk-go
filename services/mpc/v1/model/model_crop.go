package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Crop struct {
	Duration *int32 `json:"duration,omitempty"`
}

func (o Crop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Crop struct{}"
	}

	return strings.Join([]string{"Crop", string(data)}, " ")
}
