package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMediaProcessReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	TemplateId *string `json:"template_id,omitempty"`
}

func (o CreateMediaProcessReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMediaProcessReq struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessReq", string(data)}, " ")
}
