package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OutputFileInfo struct {
	OutputFileName *string `json:"output_file_name,omitempty"`

	ExecDescription *string `json:"exec_description,omitempty"`

	MetaData *SourceInfo `json:"meta_data,omitempty"`
}

func (o OutputFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputFileInfo struct{}"
	}

	return strings.Join([]string{"OutputFileInfo", string(data)}, " ")
}
