package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TranscodeDetail struct {
	MultitaskInfo *[]MultiTaskInfo `json:"multitask_info,omitempty"`

	InputFile *SourceInfo `json:"input_file,omitempty"`
}

func (o TranscodeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeDetail struct{}"
	}

	return strings.Join([]string{"TranscodeDetail", string(data)}, " ")
}
