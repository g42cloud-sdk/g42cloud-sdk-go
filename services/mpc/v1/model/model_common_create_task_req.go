package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CommonCreateTaskReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	UserData *string `json:"user_data,omitempty"`
}

func (o CommonCreateTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonCreateTaskReq struct{}"
	}

	return strings.Join([]string{"CommonCreateTaskReq", string(data)}, " ")
}
