package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
