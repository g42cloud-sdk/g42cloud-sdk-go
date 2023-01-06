package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddonTemplate struct {
	Kind string `json:"kind"`

	ApiVersion string `json:"apiVersion"`

	Metadata *Metadata `json:"metadata"`

	Spec *Templatespec `json:"spec"`
}

func (o AddonTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonTemplate struct{}"
	}

	return strings.Join([]string{"AddonTemplate", string(data)}, " ")
}
