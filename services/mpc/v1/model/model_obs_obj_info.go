package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ObsObjInfo struct {
	Bucket string `json:"bucket"`

	Location string `json:"location"`

	Object string `json:"object"`

	FileName *string `json:"file_name,omitempty"`
}

func (o ObsObjInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsObjInfo struct{}"
	}

	return strings.Join([]string{"ObsObjInfo", string(data)}, " ")
}
