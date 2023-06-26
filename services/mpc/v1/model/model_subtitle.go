package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Subtitle struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Inputs *[]MulInputFileInfo `json:"inputs,omitempty"`

	SubtitleType *int32 `json:"subtitle_type,omitempty"`
}

func (o Subtitle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subtitle struct{}"
	}

	return strings.Join([]string{"Subtitle", string(data)}, " ")
}
