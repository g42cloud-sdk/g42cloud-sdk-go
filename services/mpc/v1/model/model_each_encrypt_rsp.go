package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EachEncryptRsp struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	OutputFileName *[]string `json:"output_file_name,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o EachEncryptRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EachEncryptRsp struct{}"
	}

	return strings.Join([]string{"EachEncryptRsp", string(data)}, " ")
}
