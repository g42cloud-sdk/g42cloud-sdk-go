package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateRemuxTaskReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	OutputParam *RemuxOutputParam `json:"output_param,omitempty"`
}

func (o CreateRemuxTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskReq struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskReq", string(data)}, " ")
}
