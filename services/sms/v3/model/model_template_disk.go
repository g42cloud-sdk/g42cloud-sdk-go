package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TemplateDisk struct {
	Id *int64 `json:"id,omitempty"`

	Index int32 `json:"index"`

	Name string `json:"name"`

	Disktype string `json:"disktype"`

	Size int64 `json:"size"`

	DeviceUse *string `json:"device_use,omitempty"`
}

func (o TemplateDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateDisk struct{}"
	}

	return strings.Join([]string{"TemplateDisk", string(data)}, " ")
}
