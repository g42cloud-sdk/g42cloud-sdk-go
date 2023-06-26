package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAnimatedGraphicsTaskReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	OutputParam *AnimatedGraphicsOutputParam `json:"output_param,omitempty"`
}

func (o CreateAnimatedGraphicsTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskReq struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskReq", string(data)}, " ")
}
