package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateExtractTaskResponse struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	OutputFileName *string `json:"output_file_name,omitempty"`

	Description *string `json:"description,omitempty"`

	Metadata       *MetaData `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateExtractTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtractTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateExtractTaskResponse", string(data)}, " ")
}
