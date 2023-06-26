package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateExtractTaskReq struct {
	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Sync *int32 `json:"sync,omitempty"`
}

func (o CreateExtractTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtractTaskReq struct{}"
	}

	return strings.Join([]string{"CreateExtractTaskReq", string(data)}, " ")
}
