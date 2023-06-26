package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MbTaskParameter struct {
	StatusDescription *string `json:"status_description,omitempty"`

	OutputFilename *string `json:"output_filename,omitempty"`

	Metadata *MetaData `json:"metadata,omitempty"`
}

func (o MbTaskParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MbTaskParameter struct{}"
	}

	return strings.Join([]string{"MbTaskParameter", string(data)}, " ")
}
