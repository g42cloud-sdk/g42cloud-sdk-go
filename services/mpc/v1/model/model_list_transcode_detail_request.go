package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTranscodeDetailRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	XVodProjectId *string `json:"x-vod-projectId,omitempty"`

	TaskId *[]string `json:"task_id,omitempty"`
}

func (o ListTranscodeDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodeDetailRequest", string(data)}, " ")
}
