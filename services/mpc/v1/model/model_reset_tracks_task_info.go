package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ResetTracksTaskInfo struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	Description *string `json:"description,omitempty"`

	OutputFilename *string `json:"output_filename,omitempty"`

	TracksInfo *[]TracksInfo `json:"tracks_info,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`
}

func (o ResetTracksTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetTracksTaskInfo struct{}"
	}

	return strings.Join([]string{"ResetTracksTaskInfo", string(data)}, " ")
}
