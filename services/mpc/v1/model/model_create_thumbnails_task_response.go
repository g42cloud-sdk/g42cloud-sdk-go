package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateThumbnailsTaskResponse struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	OutputFileName *string `json:"output_file_name,omitempty"`

	ThumbnailTime *string `json:"thumbnail_time,omitempty"`

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateThumbnailsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateThumbnailsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateThumbnailsTaskResponse", string(data)}, " ")
}
