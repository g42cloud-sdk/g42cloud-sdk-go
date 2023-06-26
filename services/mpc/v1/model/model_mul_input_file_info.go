package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MulInputFileInfo struct {
	Language *string `json:"language,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`
}

func (o MulInputFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MulInputFileInfo struct{}"
	}

	return strings.Join([]string{"MulInputFileInfo", string(data)}, " ")
}
