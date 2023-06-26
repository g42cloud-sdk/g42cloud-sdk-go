package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTranscodingTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	XVodProjectId *string `json:"x-vod-projectId,omitempty"`

	Body *CreateTranscodingReq `json:"body,omitempty"`
}

func (o CreateTranscodingTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTranscodingTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTranscodingTaskRequest", string(data)}, " ")
}
