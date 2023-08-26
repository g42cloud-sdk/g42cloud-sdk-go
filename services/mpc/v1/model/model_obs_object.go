package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ObsObject struct {
	FileName *string `json:"file_name,omitempty"`

	Size *int64 `json:"size,omitempty"`

	LastModified *string `json:"last_modified,omitempty"`
}

func (o ObsObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsObject struct{}"
	}

	return strings.Join([]string{"ObsObject", string(data)}, " ")
}
