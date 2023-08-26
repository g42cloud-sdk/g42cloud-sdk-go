package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CompressRules struct {
	CompressType *string `json:"compress_type,omitempty"`

	CompressFileType *string `json:"compress_file_type,omitempty"`
}

func (o CompressRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompressRules struct{}"
	}

	return strings.Join([]string{"CompressRules", string(data)}, " ")
}
